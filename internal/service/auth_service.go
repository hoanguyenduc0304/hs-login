package service

import (
	"hs-login/internal/model"
)

type IAuthService interface {
	Login(param *model.LoginUserParam) (*model.User, error)
}
