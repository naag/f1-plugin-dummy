package support

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/naag/f1-test-utils/pkg/support/errorh"
)

type gzipTransport struct {
	underlying http.RoundTripper
}

func (g gzipTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	if request.Body != nil && strings.Contains(request.URL.RawPath, "payments") {
		var buf bytes.Buffer
		compressed := gzip.NewWriter(&buf)
		defer errorh.SafeClose(compressed)

		_, err := io.Copy(compressed, request.Body)
		if err != nil {
			return nil, err
		}
		request.Body = ioutil.NopCloser(bytes.NewReader(buf.Bytes()))
		enc := request.Header.Get("Content-Encoding")
		if enc == "" {
			enc = "gzip"
		} else {
			enc = enc + ",gzip"
		}
		request.Header.Add("Content-Encoding", enc)
	}
	return g.underlying.RoundTrip(request)

}
