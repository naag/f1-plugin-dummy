package support

import (
	"fmt"
	"strings"
	"sync"

	"github.com/form3tech/go-form3-web/httpclient"

	"github.com/form3tech-oss/f1/pkg/f1/testing"
	"github.com/form3tech/go-form3/client"
	"github.com/form3tech/go-logger/log"
	"github.com/form3tech/go-paymentapi-client/client/payments"
	rc "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

var paymentAPIClients = make(map[string]*payments.Client)
var paymentAPIClientsMtx sync.Mutex

// PaymentAPIClientForTest creates a PaymentAPI client with client config values taken from environment variables (if not overridden in t.Environment[] config).
func PaymentAPIClientForTest(t *testing.T) *payments.Client {
	targetHost, unset := GetFromTWithEnvFallback(t, FORM3_HOST, []string{})
	clientID, unset := GetFromTWithEnvFallback(t, FORM3_CLIENT_ID, unset)
	secret, unset := GetFromTWithEnvFallback(t, FORM3_CLIENT_SECRET, unset)
	organisationID, unset := GetFromTWithEnvFallback(t, FORM3_ORGANISATION_ID, unset)
	userID, unset := GetFromTWithEnvFallback(t, FORM3_USER_ID, unset)

	key := fmt.Sprintf("%v:%v:%v", clientID, secret, organisationID)
	paymentAPIClientsMtx.Lock()
	defer paymentAPIClientsMtx.Unlock()

	paymentAPIClient := paymentAPIClients[key]
	if paymentAPIClient != nil {
		return paymentAPIClient
	}

	if len(unset) > 0 {
		log.Info(fmt.Sprintf("FORM3_HOST, FORM3_CLIENT_ID, FORM3_CLIENT_SECRET, FORM3_ORGANISATION_ID and FORM3_USER_ID should have values.\n\tThe following values were blank [%s]", strings.Join(unset, ", ")))
	}

	log.Info("Building Payment API client from env vars")
	paymentAPIClient = NewPaymentAPIClientWithConfig(targetHost, clientID, userID, secret)

	paymentAPIClients[key] = paymentAPIClient
	return paymentAPIClient
}

func NewPaymentAPIClientWithConfig(host, clientID, userID, secret string) *payments.Client {
	u, err := GetURLForHost(host)
	if err != nil {
		panic(err)
	}

	c := httpclient.NewClientConfig(clientID, secret, userID, u)

	transportConfig := client.DefaultTransportConfig().
		WithHost(c.HostUrl.Host).
		WithBasePath("/v1/transaction").
		WithSchemes([]string{c.HostUrl.Scheme})

	httpClient := httpclient.NewHttpClient(c)

	rt := rc.NewWithClient(transportConfig.Host, transportConfig.BasePath, transportConfig.Schemes, httpClient)
	return payments.New(rt, strfmt.Default)
}
