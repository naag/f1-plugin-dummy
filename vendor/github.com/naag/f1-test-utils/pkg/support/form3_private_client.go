package support

import (
	"fmt"
	"strings"
	"sync"

	"github.com/form3tech/go-form3"

	"github.com/form3tech-oss/f1/pkg/f1/testing"
	"github.com/form3tech/go-form3/client"
	"github.com/form3tech/go-logger/log"
)

var f3PrivateClients = make(map[string]*form3.AuthenticatedClient)
var f3PrivateClientsMtx sync.Mutex

// F3PrivateForTest creates a Form3 private client with client config values taken from environment variables (if not overridden in t.Environment[] config).
func F3PrivateForTest(t *testing.T) *form3.AuthenticatedClient {
	targetHost, unset := GetFromTWithEnvFallback(t, FORM3_HOST, []string{})
	clientID, unset := GetFromTWithEnvFallback(t, FORM3_CLIENT_ID, unset)
	secret, unset := GetFromTWithEnvFallback(t, FORM3_CLIENT_SECRET, unset)
	organisationID, unset := GetFromTWithEnvFallback(t, FORM3_ORGANISATION_ID, unset)

	key := fmt.Sprintf("%v:%v:%v", clientID, secret, organisationID)
	f3PrivateClientsMtx.Lock()
	defer f3PrivateClientsMtx.Unlock()

	f3PrivateClient := f3PrivateClients[key]
	if f3PrivateClient != nil {
		return f3PrivateClient
	}

	if len(unset) > 0 {
		log.Info(fmt.Sprintf("FORM3_HOST, FORM3_CLIENT_ID, FORM3_CLIENT_SECRET and FORM3_ORGANISATION_ID should have values.\n\tThe following values were blank [%s]", strings.Join(unset, ", ")))
	}

	log.Info("Building Form3 private client from env vars")
	f3PrivateClient = NewForm3PrivateClientWithConfig(targetHost, clientID, secret)

	f3PrivateClients[key] = f3PrivateClient
	return f3PrivateClient
}

func NewForm3PrivateClientWithConfig(host, clientID, secret string) *form3.AuthenticatedClient {
	u, _ := GetURLForHost(host)

	httpClient := form3.NewAuthenticatedClient(&client.TransportConfig{
		BasePath: u.Path,
		Host:     u.Host,
		Schemes:  []string{u.Scheme},
	})

	if err := httpClient.Authenticate(clientID, secret); err != nil {
		panic(err.Error())
	}

	return httpClient
}

func GetFromTWithEnvFallback(t *testing.T, name string, unset []string) (string, []string) {
	val, ok := t.Environment[name]

	if !ok {
		val = getEnv(t, name)
	}

	return val, unset
}
