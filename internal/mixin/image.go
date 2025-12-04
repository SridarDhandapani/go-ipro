package mixin

import (
	"net/http"

	"github.com/SridarDhandapani/go-ipro/pkg/api"
	"github.com/SridarDhandapani/go-ipro/pkg/model"
	"github.com/google/go-querystring/query"
	"gopkg.in/validator.v2"
)

type ImageMixin struct{}

// SetImageCaptureMode set image capture configurations.
func (i *ImageMixin) SetImageCaptureMode(mode *model.CaptureMode) func(h *api.APIHandler) error {
	return func(h *api.APIHandler) error {
		if err := validator.Validate(mode); err != nil {
			return err
		}

		payload, err := query.Values(mode)
		if err != nil {
			return err
		}
		_, err = h.Request(http.MethodPost, "set_imgmode", nil, payload)
		return err
	}
}

// SetDayNightMode set day-night mode.
func (i *ImageMixin) SetDayNightMode(mode *model.DayNightMode) func(h *api.APIHandler) error {
	return func(h *api.APIHandler) error {
		if err := validator.Validate(mode); err != nil {
			return err
		}

		payload, err := query.Values(mode)
		if err != nil {
			return err
		}
		_, err = h.Request(http.MethodPost, "image_adjust", nil, payload)
		return err
	}
}
