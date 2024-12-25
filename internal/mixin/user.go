package mixin

import (
	"net/http"
	"net/url"

	"github.com/SridarDhandapani/go-ipro/pkg/api"
	ce "github.com/SridarDhandapani/go-ipro/pkg/error"
	"github.com/SridarDhandapani/go-ipro/pkg/model"
	"github.com/google/go-querystring/query"
)

type UserMixin struct{}

// RegisterAdmin register the initial admin user.
func (u *UserMixin) RegisterAdmin(admin *model.User) func(h *api.APIHandler) error {
	return func(h *api.APIHandler) error {
		cmd := "reg_admin"

		token, err := h.GetInitialRandNum()
		if err != nil {
			return err
		}
		if token == "" {
			return ce.ErrAlreadyInitialised
		}

		if err := admin.Encrypt(token); err != nil {
			return err
		}

		payload, err := query.Values(admin)
		if err != nil {
			return err
		}

		// Add additional required data
		payload.Set("ParamEnc", "1")
		payload.Set("repassword", admin.Password)
		payload.Set("Randomnum", token)

		_, err = h.Request(http.MethodPost, cmd, nil, payload)
		return err
	}
}

// RegisterUser register a user with the given name and access level.
func (u *UserMixin) RegisterUser(user *model.User) func(h *api.APIHandler) error {
	return func(h *api.APIHandler) error {
		cmd := "reg_user"

		token, err := h.GetRandParam()
		if err != nil {
			return err
		}
		if err := user.Encrypt(token); err != nil {
			return err
		}

		payload, err := query.Values(user)
		if err != nil {
			return err
		}

		// Add additional required data
		payload.Set("ParamEnc", "1")
		payload.Set("repassword", user.Password)
		payload.Set("Randomnum", token)

		_, err = h.Request(http.MethodPost, cmd, nil, payload)
		return err
	}
}

// DeleteUser deletes a user.
func (u *UserMixin) DeleteUser(name string) func(h *api.APIHandler) error {
	return func(h *api.APIHandler) error {
		cmd := "del_user"

		token, err := h.GetRandParam()
		if err != nil {
			return err
		}

		payload := url.Values{}
		payload.Set("name", name)
		payload.Set("Randomnum", token)

		_, err = h.Request(http.MethodPost, cmd, nil, payload)
		return err
	}
}
