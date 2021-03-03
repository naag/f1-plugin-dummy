package support

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/naag/f1-test-utils/pkg/support/errorh"

	"github.com/form3tech-oss/f1/pkg/f1/testing"
	log "github.com/sirupsen/logrus"
)

const DefaultCertsPath = "src/github.com/form3tech/docker-jenkins/certs"

type FpsInProxy struct {
	host   string
	client *http.Client
}

func (p *FpsInProxy) Post(ctx context.Context, path string, requestBody *strings.Reader) (resp *http.Response, body string, err error) {
	protocol := "https"
	if os.Getenv("FORM3_PROTOCOL") != "" {
		protocol = os.Getenv("FORM3_PROTOCOL")
	}
	req, _ := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s://%s%s", protocol, p.host, path), requestBody)
	resp, err = p.client.Do(req)
	if err != nil {
		return
	}
	if resp != nil {
		defer errorh.SafeClose(resp.Body)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	body = string(bodyBytes)
	return
}

func NewFpsInProxy(t *testing.T) *FpsInProxy {

	certsPath, hasPath := os.LookupEnv("FPS_IN_CERTS_PATH")
	if !hasPath {
		certsPath = path.Join(os.Getenv("GOPATH"), DefaultCertsPath)
	}

	certFile := path.Join(certsPath, "client.cert.pem")
	keyFile := path.Join(certsPath, "client.key.pem")

	for _, f := range []string{certFile, keyFile} {
		if _, certError := os.Stat(f); os.IsNotExist(certError) {
			log.Infof("Cannot find cert file %s", f)
			t.Errorf(
				"unable to find %s for fps-in in %s. "+
					"This should be in either $FPS_IN_CERTS_PATH or $GOPATH/%s",
				path.Base(f),
				certsPath,
				DefaultCertsPath,
			)
		}
	}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	t.Require.Nil(err)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}
	transport := &http.Transport{TLSClientConfig: tlsConfig, Proxy: http.ProxyFromEnvironment}

	if os.Getenv("FPS_IN_HOST") == "" {
		t.Log.Fatal("Missing environment variable FPS_IN_HOST for admissions test")
	}
	return &FpsInProxy{
		host:   os.Getenv("FPS_IN_HOST"),
		client: &http.Client{Transport: transport},
	}
}
