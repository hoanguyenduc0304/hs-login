package service

import (
	"hs-login/internal/model"
	"hs-login/internal/repository"
	"sync"
)

var (
	authOnce    sync.Once
	authService IAuthService
)

func NewAuthService(authRepo repository.IAuthRepository) IAuthService {
	authOnce.Do(func() {
		authService = &AuthService{
			authRepository: authRepo,
		}
	})
	return authService
}

type AuthService struct {
	authRepository repository.IAuthRepository
}

func (a *AuthService) Login(auth *model.LoginUserParam) (*model.User, error) {
	return a.authRepository.CheckUsername(auth.Username)
}
