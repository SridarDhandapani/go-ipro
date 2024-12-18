package camera

import (
	"github.com/SridarDhandapani/go-ipro/internal/api"
	"github.com/SridarDhandapani/go-ipro/internal/camera"
)

type Camera struct {
	*camera.CameraHandler
}

type options struct {
	name     string
	username string
	password string
	host     string
	apiOpts  []api.ApiHandlerOption
}

type CameraOption interface {
	apply(*options)
}

// options.name
type nameOption string

func (o nameOption) apply(opts *options) {
	opts.name = string(o)
}
func WithName(name string) CameraOption {
	return nameOption(name)
}

// options.username
type usernameOption string

func (o usernameOption) apply(opts *options) {
	opts.username = string(o)
}
func WithUsername(username string) CameraOption {
	return usernameOption(username)
}

// options.password
type passwordOption string

func (o passwordOption) apply(opts *options) {
	opts.password = string(o)
}
func WithPassword(password string) CameraOption {
	return passwordOption(password)
}

// options.host
type hostOption string

func (o hostOption) apply(opts *options) {
	opts.host = string(o)
}
func WithHost(host string) CameraOption {
	return hostOption(host)
}

// options.apiOpts
type apiOptsOption struct {
	opts []api.ApiHandlerOption
}

func (o apiOptsOption) apply(opts *options) {
	opts.apiOpts = o.opts
}
func WithApiOpts(opts []api.ApiHandlerOption) CameraOption {
	return apiOptsOption{opts}
}

func NewCamera(name string, opts ...CameraOption) (c *Camera, err error) {
	options := options{
		name:     "Reolink",
		username: "admin",
		password: "",
		apiOpts:  nil,
	}
	for _, opt := range opts {
		opt.apply(&options)
	}

	cameraHandler, err := camera.NewCameraHandler(options.username, options.password, options.apiOpts...)
	if err != nil {
		return nil, err
	}
	c = &Camera{cameraHandler}
	return c, nil
}
