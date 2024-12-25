package mixin

import (
	"errors"
	"net/http"
	"net/url"
	"slices"

	"github.com/SridarDhandapani/go-ipro/internal/utils"
	"github.com/SridarDhandapani/go-ipro/pkg/api"
	"github.com/SridarDhandapani/go-ipro/pkg/model"
	"github.com/google/go-querystring/query"
	"gopkg.in/validator.v2"
)

type SystemMixin struct{}

// GetInfo gets camera info.
func (s *SystemMixin) GetInfo() func(h *api.APIHandler) (*model.CameraInfo, error) {
	return func(h *api.APIHandler) (*model.CameraInfo, error) {
		params := url.Values{}
		params.Set("FILE", "1")
		resp, err := h.Request(http.MethodGet, "getinfo", params, nil)
		if err != nil {
			return nil, err
		}

		var data *model.CameraInfo
		err = utils.MapToStruct(resp, &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}

// GetData gets the camera configuration data.
func (s *SystemMixin) GetData() func(h *api.APIHandler) (*model.CameraData, error) {
	return func(h *api.APIHandler) (*model.CameraData, error) {
		resp, err := h.Request(http.MethodGet, "getdata", nil, nil)
		if err != nil {
			return nil, err
		}
		var data *model.CameraData
		err = utils.MapToStruct(resp, &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}

// SetData updates camera configuration.
func (s *SystemMixin) SetData(v interface{}) func(h *api.APIHandler) error {
	return func(h *api.APIHandler) error {
		if err := validator.Validate(v); err != nil {
			return err
		}

		payload, err := query.Values(v)
		if err != nil {
			return err
		}
		_, err = h.Request(http.MethodPost, "setdata", nil, payload)
		return err
	}
}

func (s *SystemMixin) Reset(mode string) func(h *api.APIHandler) error {
	return func(h *api.APIHandler) error {
		if mode == "" {
			mode = "all"
		}
		rp, err := h.GetRandParam()
		if err != nil {
			return err
		}

		validCmd := []string{"all", "data", "html", "reset", "restart"}
		if !(slices.Contains(validCmd, mode)) {
			return errors.New("reset mode not supported")
		}

		payload := url.Values{}
		payload.Set("cmd", mode)
		payload.Set("Randomnum", rp)

		_, err = h.Request(http.MethodPost, "initial", nil, payload)
		return err
	}
}
