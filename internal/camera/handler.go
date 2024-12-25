package camera

import (
	"github.com/SridarDhandapani/go-ipro/internal/mixin"
	"github.com/SridarDhandapani/go-ipro/pkg/api"
)

type CameraHandler struct {
	*mixin.ImageMixin
	*mixin.StreamMixin
	*mixin.SystemMixin
	*mixin.UserMixin
	*api.APIHandler
}

type options struct {
	name string
}

type CameraOptions interface {
	apply(*options)
}

func NewCameraHandler(username, password string, apiOpts ...api.APIHandlerOption) (*CameraHandler, error) {
	apiHandler := api.NewAPIHandler(username, password, apiOpts...)
	return &CameraHandler{
		&mixin.ImageMixin{},
		&mixin.StreamMixin{},
		&mixin.SystemMixin{},
		&mixin.UserMixin{},
		apiHandler,
	}, nil
}
