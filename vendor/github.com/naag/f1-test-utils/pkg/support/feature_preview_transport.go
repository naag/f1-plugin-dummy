package support

import (
	"net/http"
	"net/url"
	"strings"
)

type PreviewConfig struct {
	ReadModel bool
}

// featurePreviewTransport toggles certain pre-defined feature switches.
type featurePreviewTransport struct {
	underlying http.RoundTripper
	PreviewConfig
}

func (t featurePreviewTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	if t.ReadModel && strings.HasPrefix(request.URL.Path, "/v1/transaction/payments") && request.Method == http.MethodGet {
		query, _ := url.ParseQuery(request.URL.RawQuery)
		query.Add("source", "read-model")
		request.URL.RawQuery = query.Encode()
	}
	return t.underlying.RoundTrip(request)
}
