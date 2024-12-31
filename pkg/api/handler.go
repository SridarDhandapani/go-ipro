package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/SridarDhandapani/go-ipro/internal/utils"
	"github.com/icholy/digest"
)

var headers = map[string][]string{
	"Content-Type": {"application/x-www-form-urlencoded"},
}

type APIHandler struct {
	*options
}

type options struct {
	client         *http.Client
	scheme         Scheme
	host           string
	port           int
	endpointPrefix string
}

type APIHandlerOption interface {
	apply(*options)
}

// options.client
type clientOption struct {
	client *http.Client
}

func (o clientOption) apply(opts *options) {
	opts.client = o.client
}

func WithClient(client *http.Client) APIHandlerOption {
	return clientOption{client: client}
}

// options.scheme
type schemeOption struct {
	Scheme
}

func (o schemeOption) apply(opts *options) {
	opts.scheme = o.Scheme
}

func WithScheme(scheme Scheme) APIHandlerOption {
	return schemeOption{scheme}
}

// options.host
type hostOption string

func (o hostOption) apply(opts *options) {
	opts.host = string(o)
}

func WithHost(host string) APIHandlerOption {
	return hostOption(host)
}

// options.port
type portOption int

func (o portOption) apply(opts *options) {
	opts.port = int(o)
}

func WithPort(port int) APIHandlerOption {
	return portOption(port)
}

// options.endpointPrefix
type endpointPrefixOption string

func (o endpointPrefixOption) apply(opts *options) {
	opts.endpointPrefix = string(o)
}

func WithEndpointPrefix(basePath string) APIHandlerOption {
	return endpointPrefixOption(basePath)
}

func NewAPIHandler(username, password string, opts ...APIHandlerOption) *APIHandler {
	options := &options{
		client:         &http.Client{},
		scheme:         HTTPS,
		host:           "",
		port:           443,
		endpointPrefix: "/cgi-bin",
	}
	for _, opt := range opts {
		opt.apply(options)
	}

	client := &http.Client{Transport: &digest.Transport{
		Username:  username,
		Password:  password,
		Transport: options.client.Transport,
	}}
	options.client = client

	return &APIHandler{
		options,
	}
}

func (ah *APIHandler) Request(method string, command string, params url.Values, payload url.Values) (map[string]interface{}, error) {
	urlPath, _ := url.JoinPath(ah.endpointPrefix, command)

	resp, err := ah.RequestRaw(method, urlPath, params, payload)
	if err != nil {
		return nil, err
	}
	return utils.RespToMap(string(resp)), nil
}

func (ah *APIHandler) RequestRaw(method string, urlPath string, params url.Values, payload url.Values) ([]byte, error) {
	urlConcat := fmt.Sprintf("%s://%s:%d%s?%s", ah.scheme.String(), ah.host, ah.port, urlPath, params.Encode())

	reqURL, err := url.Parse(urlConcat)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, reqURL.String(), strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header = headers

	resp, err := ah.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint:errcheck
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed, err: %v", resp.Status)
	}

	var body bytes.Buffer
	_, err = io.Copy(&body, resp.Body)
	if err != nil {
		return nil, err
	}

	return body.Bytes(), nil
}

// GetInitialRandNum fetches the initial random number after the camera is reset to default.
func (ah *APIHandler) GetInitialRandNum() (string, error) {
	urlPath := "/js/crandomnum4init.js"
	resp, err := ah.RequestRaw(http.MethodGet, urlPath, nil, nil)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`var (.+)=([^\r\n]+)`)
	respMap := utils.RespToMapWithRegex(string(resp), re)

	var data struct {
		RandNum string `json:"randomnum"`
	}
	err = utils.MapToStruct(respMap, &data)
	if err != nil {
		return "", err
	}
	return data.RandNum, nil
}

// GetRandParam fetches the random param to validate requests.
func (ah *APIHandler) GetRandParam() (string, error) {
	urlPath := "/admin/setup_ini.html"
	resp, err := ah.RequestRaw(http.MethodGet, urlPath, nil, nil)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`var (.+)=([^\r\n]+);`)
	respMap := utils.RespToMapWithRegex(string(resp), re)

	var data struct {
		RandParam string `json:"giRandParam"`
	}
	err = utils.MapToStruct(respMap, &data)
	if err != nil {
		return "", err
	}
	return data.RandParam, nil
}
