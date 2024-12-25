package camera

import (
	"github.com/SridarDhandapani/go-ipro/internal/camera"
	"github.com/SridarDhandapani/go-ipro/pkg/api"
)

type Camera struct {
	*camera.CameraHandler
}

type options struct {
	name     string
	username string
	password string
	apiOpts  []api.APIHandlerOption
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

// options.apiOpts
type apiOptsOption struct {
	opts []api.APIHandlerOption
}

func (o apiOptsOption) apply(opts *options) {
	opts.apiOpts = o.opts
}

func WithAPIOpts(opts ...api.APIHandlerOption) CameraOption {
	return apiOptsOption{opts}
}

func NewCamera(opts ...CameraOption) (c *Camera, err error) {
	options := options{
		name:     "iPro",
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
