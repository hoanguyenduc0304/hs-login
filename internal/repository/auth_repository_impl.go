package repository

import (
	"hs-login/internal/model"
	"hs-login/server"
	"sync"
)

var (
	authOnce       sync.Once
	authRepository IAuthRepository
)

func NewAuthRepository(database server.Database) IAuthRepository {
	authOnce.Do(func() {
		authRepository = &AuthRepositoryImpl{Database: database}
	})
	return authRepository
}

type AuthRepositoryImpl struct {
	Database server.Database
}

func (a *AuthRepositoryImpl) CheckUsername(uname string) (*model.User, error) {
	var id int64
	var username, password, name string
	row := a.Database.Db.QueryRow(`SELECT id, name, username, password FROM "user" WHERE username = $1`, uname)
	err := row.Scan(&id, &name, &username, &password)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Id:       id,
		Name:     name,
		Username: username,
		Password: password,
	}, nil
}
