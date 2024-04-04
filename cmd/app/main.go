package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"hs-login/internal/api"
	"hs-login/internal/repository"
	"hs-login/internal/service"
	"hs-login/server"
	"net/http"
)

var db server.Database

func main() {
	defer db.Close()
	fmt.Println("Starting server")
	err := db.InitPSQL()
	if err != nil {
		return
	}

	// repositories
	authRepo := repository.NewAuthRepository(db)
	// services
	authService := service.NewAuthService(authRepo)
	//controller
	c := api.New(authService)

	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Post("/login", c.Login)
	http.ListenAndServe(":3000", r)
}
