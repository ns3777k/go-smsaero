package smsaero

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"fmt"
	"net/url"
)

const (
	testClientLogin = "my@login.com"
	testClientPassword = "098f6bcd4621d373cade4e832627b4f6"
	stubsDir = "stubs"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func assertRequiredFields(t *testing.T, form url.Values) {
	assert.Equal(t, testClientLogin, form.Get("user"))
	assert.Equal(t, testClientPassword, form.Get("password"))
	assert.Equal(t, responseType, form.Get("answer"))
}

func getStub(t *testing.T, stubName string) []byte {
	stubPath := fmt.Sprintf("%s/%s.json", stubsDir, stubName)
	content, err := ioutil.ReadFile(stubPath)
	if err != nil {
		t.Errorf("getStub error %v", err)
	}

	return content
}

func setUpTestServe() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewClient(nil, testClientLogin, testClientPassword)
	client.BaseURL = server.URL
}

func tearDownTestServe() {
	server.Close()
}

func TestNewClient(t *testing.T) {
	client := NewClient(nil, "test", "test")
	assert.Equal(t, "test", client.username)
	assert.Equal(t, "test", client.password)
	assert.Equal(t, client.client, http.DefaultClient)
}

func TestNewClient_customClient(t *testing.T) {
	transport := &http.Transport{}
	httpClient := &http.Client{Transport: transport}

	client := NewClient(httpClient, "test", "test")
	assert.Equal(t, "test", client.username)
	assert.Equal(t, "test", client.password)
	assert.Equal(t, client.client, httpClient)
}
