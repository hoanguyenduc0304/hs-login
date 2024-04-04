package model

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username" `
	Password string `json:"password"`
}

type LoginUserParam struct {
	Username string `json:"username" validate:"required,max=30"`
	Password string `json:"password" validate:"required,min=8,max=30"`
}
