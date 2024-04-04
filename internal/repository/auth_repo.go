package repository

import (
	"hs-login/internal/model"
)

type IAuthRepository interface {
	CheckUsername(username string) (*model.User, error)
}
