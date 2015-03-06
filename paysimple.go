package paysimple

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type api struct {
	baseURL url.URL
	config  Config
	backend backend

	// Resource endpoints
	Payments  *Payments
	Customers *Customers
}

func (api *api) decodeError(resp *http.Response) error {
	var empty Empty
	if err := json.NewDecoder(resp.Body).Decode(&empty); err != nil {
		return fmt.Errorf("paysimple: failed to decode error: %s", err)
	}
	return empty.Meta.Errors
}

func (api *api) request(method string, uri *url.URL, body io.Reader) (*http.Request, error) {
	// TODO No body for now
	req, err := http.NewRequest(method, uri.String(), body)
	if err != nil {
		return nil, err
	}

	// Add the authorization header
	auth := CreateAuthorization(api.config.Username, api.config.SecretKey)
	req.Header.Set("Authorization", auth)
	return req, nil
}

func (api *api) Get(uri *url.URL) (*http.Request, error) {
	return api.request("GET", uri, nil)
}

func (api *api) Post(uri *url.URL, v interface{}) (*http.Request, error) {
	content, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return api.request("POST", uri, bytes.NewBuffer(content))
}

func (api *api) URL(path ...string) *url.URL {
	api.baseURL.Path = strings.Join(path, "/")
	return &api.baseURL
}

type Config struct {
	Username, SecretKey string
}

func (config Config) IsValid() bool {
	// TODO better validation - expected length / prefix?
	return config.Username != "" && config.SecretKey != ""
}

func Env() (config Config) {
	config.Username = os.Getenv("PAYSIMPLE_USER")
	config.SecretKey = os.Getenv("PAYSIMPLE_SECRET")
	return
}

func create(baseURL url.URL) *api {
	config := Env()
	if !config.IsValid() {
		panic("Failed to parse environmental variables - are they set?")
	}
	api := &api{
		baseURL: baseURL,
		config:  config,
		backend: &http.Client{},
	}

	api.Payments = &Payments{api: api}
	api.Customers = &Customers{api: api}
	return api
}

func API() *api {
	return create(url.URL{Scheme: "https", Host: "api.paysimple.com"})
}

func Sandbox() *api {
	return create(url.URL{Scheme: "https", Host: "sandbox-api.paysimple.com"})
}