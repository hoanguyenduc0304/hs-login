package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"hs-login/internal/model"
	"hs-login/internal/security"
	"hs-login/internal/service"
	"net/http"
)

type AuthApi struct {
	authService service.IAuthService
}

func NewAuthHandler(authService service.IAuthService) AuthApi {
	return AuthApi{
		authService: authService,
	}
}

type LoginResponse struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (l LoginResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (c *AuthApi) Login(w http.ResponseWriter, r *http.Request) {
	req := model.LoginUserParam{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		render.Render(w, r, &LoginResponse{
			Success:    false,
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	if len(req.Username) == 0 {
		render.Render(w, r, &LoginResponse{
			Success:    false,
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	v := validator.New()
	err = v.Struct(req)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {
			render.Render(w, r, &LoginResponse{
				Success:    false,
				Message:    fmt.Sprintf("%s invalid", e.Field()),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
	}

	user, err := c.authService.Login(&req)
	if err != nil {
		render.Render(w, r, &LoginResponse{
			Success:    false,
			Message:    "Wrong username or password",
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	if user == nil {
		render.Render(w, r, &LoginResponse{
			Success:    false,
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}
	if ok := security.CheckPasswordHash(req.Password, user.Password); ok {
		render.Render(w, r, &LoginResponse{
			Success:    true,
			Message:    "Login successfully",
			StatusCode: http.StatusOK,
		})
		return
	}
	render.Render(w, r, &LoginResponse{
		Success:    false,
		Message:    "Wrong username or password",
		StatusCode: http.StatusBadRequest,
	})
}
