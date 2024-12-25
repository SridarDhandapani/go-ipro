package mixin

import (
	"fmt"
	"net/http"

	"github.com/SridarDhandapani/go-ipro/pkg/api"
	"github.com/SridarDhandapani/go-ipro/pkg/model"
	"github.com/google/go-querystring/query"
	"gopkg.in/validator.v2"
)

type StreamMixin struct{}

// SetStreamConfig configures a h264 stream.
func (s *StreamMixin) SetStreamConfig(id int, stream *model.Stream) func(h *api.APIHandler) error {
	return func(h *api.APIHandler) error {
		if id < 1 || id > 4 {
			return fmt.Errorf("invalid id %d", id)
		}

		cmd := "set_h264"
		if id > 1 {
			cmd = fmt.Sprintf("set_h264_%d", id)
		}

		if err := validator.Validate(stream); err != nil {
			return err
		}

		payload, err := query.Values(stream)
		if err != nil {
			return err
		}
		_, err = h.Request(http.MethodPost, cmd, nil, payload)
		return err
	}
}
