package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/icholy/digest"
)

type ApiHandler struct {
	*options
}

type options struct {
	client         *http.Client
	scheme         Scheme
	host           string
	port           int
	endpointPrefix string
}

type ApiHandlerOption interface {
	apply(*options)
}

// options.client
type clientOption struct {
	client *http.Client
}

func (o clientOption) apply(opts *options) {
	opts.client = o.client
}
func WithClient(client *http.Client) ApiHandlerOption {
	return clientOption{client: client}
}

// options.scheme
type schemeOption struct {
	Scheme
}

func (o schemeOption) apply(opts *options) {
	opts.scheme = o.Scheme
}
func WithScheme(scheme Scheme) ApiHandlerOption {
	return schemeOption{scheme}
}

// options.host
type hostOption string

func (o hostOption) apply(opts *options) {
	opts.host = string(o)
}
func WithHost(host string) ApiHandlerOption {
	return hostOption(host)
}

// options.port
type portOption int

func (o portOption) apply(opts *options) {
	opts.port = int(o)
}
func WithPort(port int) ApiHandlerOption {
	return portOption(port)
}

// options.endpointPrefix
type endpointPrefixOption string

func (o endpointPrefixOption) apply(opts *options) {
	opts.endpointPrefix = string(o)
}
func WithEndpointPrefix(basePath string) ApiHandlerOption {
	return endpointPrefixOption(basePath)
}

func NewApiHandler(username, password string, opts ...ApiHandlerOption) *ApiHandler {
	options := &options{
		client:         nil,
		scheme:         HTTP,
		host:           "",
		port:           80,
		endpointPrefix: "/cgi-bin",
	}
	for _, opt := range opts {
		opt.apply(options)
	}

	client := &http.Client{Transport: &digest.Transport{
		Username: username,
		Password: password,
	}}
	options.client = client

	return &ApiHandler{
		options,
	}
}

func (ah *ApiHandler) Request(method, path string, payload interface{}, params url.Values) ([]byte, error) {
	urlPath, _ := url.JoinPath(ah.endpointPrefix, path)
	urlConcat := fmt.Sprintf("%s://%s:%d%s?%s", ah.scheme.String(), ah.host, ah.port, urlPath, params.Encode())

	data, err := json.Marshal([]interface{}{payload})
	if err != nil {
		return nil, err
	}

	reqUrl, err := url.Parse(urlConcat)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, reqUrl.String(), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	resp, err := ah.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
