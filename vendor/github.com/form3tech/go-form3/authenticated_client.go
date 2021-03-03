package form3

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"sync"

	"time"

	"github.com/form3tech/go-form3/client"
	"github.com/go-openapi/runtime"
	rc "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/logging"
)

var tokenCache = sync.Map{}

type CachedToken struct {
	Token   string
	Expires time.Time
}

type AuthenticatedClient struct {
	AccessToken           string
	BillingClient         *client.Form3CorelibDataStructures
	SecurityClient        *client.Form3CorelibDataStructures
	NotificationClient    *client.Form3CorelibDataStructures
	Config                *client.TransportConfig
	HttpClient            *http.Client
	OrganisationId        string
	OrganisationClient    *client.Form3CorelibDataStructures
	AssociationClient     *client.Form3CorelibDataStructures
	AccountClient         *client.Form3CorelibDataStructures
	LimitsClient          *client.Form3CorelibDataStructures
	PaymentdefaultsClient *client.Form3CorelibDataStructures
	TransactionClient     *client.Form3CorelibDataStructures
	SystemClient          *client.Form3CorelibDataStructures
}

type AuthenticatedClientCheckRedirect struct {
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req)
}

func (r *AuthenticatedClientCheckRedirect) CheckRedirect(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", via[0].Header.Get("Authorization"))
	return nil
}

func NewAuthenticatedClient(config *client.TransportConfig) *AuthenticatedClient {
	a := &AuthenticatedClientCheckRedirect{}
	var authClient *AuthenticatedClient

	h := &http.Client{
		Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			if len(authClient.AccessToken) > 0 {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authClient.AccessToken))
			}

			if logging.IsDebugOrHigher() {
				dump, errDump := httputil.DumpRequestOut(req, true)
				if errDump != nil {
					log.Fatal(errDump)
				}

				log.Printf("[DEBUG] %s %s", req.URL.String(), string(dump))
			}

			resp, err := http.DefaultTransport.RoundTrip(req)

			if err != nil {
				return nil, err
			}

			if logging.IsDebugOrHigher() {
				dump, errDump := httputil.DumpResponse(resp, true)
				if errDump != nil {
					log.Fatal(errDump)
				}

				log.Printf("[DEBUG] %s %s", req.URL.String(), string(dump))
			}

			return resp, err
		}),
		CheckRedirect: a.CheckRedirect,
	}

	rt1 := rc.NewWithClient(config.Host, "/v1/security", config.Schemes, h)
	securityClient := client.New(rt1, strfmt.Default)

	rt2 := rc.NewWithClient(config.Host, "/v1/notification", config.Schemes, h)
	notificationClient := client.New(rt2, strfmt.Default)

	rt3 := rc.NewWithClient(config.Host, "/v1/organisation", config.Schemes, h)
	organisationClient := client.New(rt3, strfmt.Default)

	rt4 := rc.NewWithClient(config.Host, "/v1/organisation/units/associations", config.Schemes, h)
	associationsClient := client.New(rt4, strfmt.Default)

	rt5 := rc.NewWithClient(config.Host, "/v1/organisation", config.Schemes, h)
	accountClient := client.New(rt5, strfmt.Default)

	rt6 := rc.NewWithClient(config.Host, "/v1/organisation/units/", config.Schemes, h)
	limitsClient := client.New(rt6, strfmt.Default)

	rt7 := rc.NewWithClient(config.Host, "/v1/transaction", config.Schemes, h)
	transactionClient := client.New(rt7, strfmt.Default)

	rt8 := rc.NewWithClient(config.Host, "/v1/organisation/units/", config.Schemes, h)
	paymentdefaultsClient := client.New(rt8, strfmt.Default)

	rt9 := rc.NewWithClient(config.Host, "/v1/system/", config.Schemes, h)
	systemClient := client.New(rt9, strfmt.Default)

    rt10 := rc.NewWithClient(config.Host, "/v1/organisation/billing/", config.Schemes, h)
    billingClient := client.New(rt10, strfmt.Default)

	authClient = &AuthenticatedClient{
		AssociationClient:     associationsClient,
		BillingClient:         billingClient,
		SecurityClient:        securityClient,
		NotificationClient:    notificationClient,
		OrganisationClient:    organisationClient,
		AccountClient:         accountClient,
		LimitsClient:          limitsClient,
		PaymentdefaultsClient: paymentdefaultsClient,
		TransactionClient:     transactionClient,
		SystemClient:          systemClient,
		HttpClient:            h,
		Config:                config,
	}

	configureRuntime(rt1, authClient)
	configureRuntime(rt2, authClient)
	configureRuntime(rt3, authClient)
	configureRuntime(rt4, authClient)
	configureRuntime(rt5, authClient)
	configureRuntime(rt6, authClient)
	configureRuntime(rt7, authClient)
	configureRuntime(rt8, authClient)
	configureRuntime(rt9, authClient)
	configureRuntime(rt10, authClient)

	return authClient
}
func configureRuntime(rt *rc.Runtime, authClient *AuthenticatedClient) {
	rt.Consumers["application/vnd.api+json;charset=UTF-8"] = runtime.JSONConsumer()
	rt.Consumers["application/vnd.api+json"] = runtime.JSONConsumer()
}

func (r *AuthenticatedClient) Authenticate(clientId string, clientSecret string) error {
	token, cached := tokenCache.Load(clientId)
	if cached && time.Now().Before(token.(CachedToken).Expires) {
		r.AccessToken = token.(CachedToken).Token
		return nil
	}

	req, _ := http.NewRequest("POST", "/v1/oauth2/token", bytes.NewBufferString("grant_type=client_credentials"))
	req.URL.Host = r.Config.Host
	req.URL.Scheme = r.Config.Schemes[0]

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(clientId+":"+clientSecret)))

	if logging.IsDebugOrHigher() {
		dump, errDump := httputil.DumpRequestOut(req, true)
		if errDump != nil {
			log.Fatal(errDump)
		}

		log.Printf("[DEBUG] %s %s", req.URL.String(), string(dump))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("error returned while authenticating, response code was %v", resp.StatusCode)
		return err
	}

	defer resp.Body.Close()

	if logging.IsDebugOrHigher() {
		dump, errDump := httputil.DumpResponse(resp, true)
		if errDump != nil {
			log.Fatal(errDump)
		}

		log.Printf("[DEBUG] %s %s", req.URL.String(), string(dump))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	loginResponse, err := getLoginResponse([]byte(body))

	if err != nil {
		return err
	}

	r.AccessToken = loginResponse.AccessToken
	tokenCache.Store(clientId, CachedToken{Token: r.AccessToken, Expires: time.Now().Add(time.Duration(loginResponse.ExpiresIn/2) * time.Second)})

	return nil
}

func getLoginResponse(body []byte) (*LoginResponse, error) {
	var s = new(LoginResponse)
	err := json.Unmarshal(body, &s)
	return s, err
}

type LoginResponse struct {
	TokenType   string `json:"token_type,omitempty"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
