package support

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"

	"github.com/form3tech-oss/f1/pkg/f1/testing"
	"github.com/form3tech-oss/go-form3/v2/pkg/form3"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client"
	"github.com/form3tech/go-logger/log"
	rc "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

var f3Clients = make(map[string]*form3.F3)
var f3ClientsMtx sync.Mutex

type Form3Credentials struct {
	ClientId       string
	ClientSecret   string
	OrganisationId string
}

const (
	FORM3_HOST            = "FORM3_HOST"
	FORM3_CLIENT_ID       = "FORM3_CLIENT_ID"
	FORM3_CLIENT_SECRET   = "FORM3_CLIENT_SECRET"
	FORM3_ORGANISATION_ID = "FORM3_ORGANISATION_ID"
	FORM3_USER_ID         = "FORM3_USER_ID"
)

func NewForm3Credentials(t *testing.T) Form3Credentials {
	clientId := t.Environment[FORM3_CLIENT_ID]
	secret := t.Environment[FORM3_CLIENT_SECRET]
	organisationId := t.Environment[FORM3_ORGANISATION_ID]

	if clientId == "" || secret == "" || organisationId == "" {
		return Form3Credentials{
			ClientId:       os.Getenv(FORM3_CLIENT_ID),
			ClientSecret:   os.Getenv(FORM3_CLIENT_SECRET),
			OrganisationId: os.Getenv(FORM3_ORGANISATION_ID),
		}
	} else {
		return Form3Credentials{
			ClientId:       clientId,
			ClientSecret:   secret,
			OrganisationId: organisationId,
		}
	}
}

// F3ForTest creates a Form3 client with client config values taken from environment variables.
func F3ForTest(t *testing.T) *form3.F3 {
	return F3ForFeaturePreviewTest(t, PreviewConfig{})
}

// F3ForFeaturePreviewTest creates a Form3 client with client config values taken from environment variables.
// Use `PreviewConfig` to enable feature previews.
func F3ForFeaturePreviewTest(t *testing.T, previewConfig PreviewConfig) *form3.F3 {
	clientID := t.Environment[FORM3_CLIENT_ID]
	secret := t.Environment[FORM3_CLIENT_SECRET]
	organisationID := t.Environment[FORM3_ORGANISATION_ID]

	key := fmt.Sprintf("%v:%v:%v", clientID, secret, organisationID)
	f3ClientsMtx.Lock()
	defer f3ClientsMtx.Unlock()

	f3Client := f3Clients[key]
	if f3Client != nil {
		return f3Client
	}

	if clientID == "" || secret == "" || organisationID == "" {
		log.Info("Building Form3 client from environment variables")
		f3Client = NewForm3ClientWithFeaturePreview(t, previewConfig)
	} else {
		log.Info("Building Form3 client from URL parameters")
		f3Client = NewForm3ClientWithConfig(os.Getenv("FORM3_HOST"), clientID, secret, organisationID, os.Getenv("FORM3_GZIP"), previewConfig)
	}

	// start with a get request to work around authentication retry limitation for post requests.
	_, _ = f3Client.Payments.GetPayment().Do()
	f3Clients[key] = f3Client
	return f3Client
}

func getEnv(t *testing.T, name string) string {
	env, ok := os.LookupEnv(name)
	t.Require.True(ok, "Missing required environment variable "+name)
	return env
}

func NewForm3Client(t *testing.T) *form3.F3 {
	return NewForm3ClientWithFeaturePreview(t, PreviewConfig{})
}

// NewForm3ClientWithFeaturePreview creates a new Form3Client with config values taken from environment variables.
// `previewFeature` configures a RoundTripper to modify requests to enable feature previews.
func NewForm3ClientWithFeaturePreview(t *testing.T, previewConfig PreviewConfig) *form3.F3 {
	return NewForm3ClientWithConfig(
		getEnv(t, "FORM3_HOST"),
		getEnv(t, "FORM3_CLIENT_ID"),
		getEnv(t, "FORM3_CLIENT_SECRET"),
		getEnv(t, "FORM3_ORGANISATION_ID"),
		os.Getenv("FORM3_GZIP"),
		previewConfig,
	)
}

func NewForm3ClientWithConfig(host, clientID, secret, organisationID, gzip string, previewConfig PreviewConfig) *form3.F3 {
	u, _ := GetURLForHost(host)

	httpclientconf := form3.NewTokenBasedClientConfig(clientID, secret, u).
		WithUnderlyingTransport(getRoundTripper(gzip, previewConfig))
	httpClient := form3.NewTokenBasedHTTPClient(httpclientconf)

	rt := rc.NewWithClient(u.Host, "/v1", []string{u.Scheme}, httpClient)
	rt.Context = context.Background()

	debug, hasDebug := os.LookupEnv("HTTP_DEBUG")
	rt.SetDebug(hasDebug && debug == "1")

	defaults := form3.NewClientDefaults()
	orgId := strfmt.UUID(organisationID)
	defaults.OrganisationId = &orgId
	cli := client.New(rt, strfmt.Default, defaults)

	return &form3.F3{
		Form3PublicAPI: *cli,
		Defaults:       defaults,
	}
}

func GetURLForHost(host string) (*url.URL, error) {
	var protocol string
	if os.Getenv("FORM3_PROTOCOL") == "" {
		protocol = "https"
	} else {
		protocol = os.Getenv("FORM3_PROTOCOL")
	}
	return url.Parse(fmt.Sprintf("%s://%s", protocol, host))
}

func getRoundTripper(gzip string, config PreviewConfig) http.RoundTripper {
	transport := http.DefaultTransport
	if b, err := strconv.ParseBool(gzip); err == nil && b {
		transport = gzipTransport{underlying: transport}
	}

	transport = featurePreviewTransport{
		underlying:    transport,
		PreviewConfig: config,
	}
	return transport
}
