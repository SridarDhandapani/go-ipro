package mixin

import (
	"net/http"

	"github.com/SridarDhandapani/go-ipro/pkg/api"
	"github.com/google/go-querystring/query"
	"gopkg.in/validator.v2"
)

type ImageMixin struct{}

// SetImageCaptureMode set image capture configurations.
func (i *ImageMixin) SetImageCaptureMode(v interface{}) func(h *api.APIHandler) error {
	return func(h *api.APIHandler) error {
		if err := validator.Validate(v); err != nil {
			return err
		}

		payload, err := query.Values(v)
		if err != nil {
			return err
		}
		_, err = h.Request(http.MethodPost, "set_imgmode", nil, payload)
		return err
	}
}

// SetDayNightMode set day-night mode.
func (i *ImageMixin) SetDayNightMode(v interface{}) func(h *api.APIHandler) error {
	return func(h *api.APIHandler) error {
		if err := validator.Validate(v); err != nil {
			return err
		}

		payload, err := query.Values(v)
		if err != nil {
			return err
		}
		_, err = h.Request(http.MethodPost, "image_adjust", nil, payload)
		return err
	}
}
