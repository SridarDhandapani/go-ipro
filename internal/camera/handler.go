package camera

import (
	"github.com/SridarDhandapani/go-ipro/internal/api"
	"github.com/SridarDhandapani/go-ipro/internal/mixin"
)

type CameraHandler struct {
	*mixin.SystemMixin
	*mixin.UserMixin
	*api.ApiHandler
}

type options struct {
	name string
}

type CameraOptions interface {
	apply(*options)
}

func NewCameraHandler(username, password string, apiOpts ...api.ApiHandlerOption) (*CameraHandler, error) {
	apiHandler := api.NewApiHandler(username, password, apiOpts...)
	return &CameraHandler{
		&mixin.SystemMixin{},
		&mixin.UserMixin{},
		apiHandler,
	}, nil
}
