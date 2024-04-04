package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"hs-login/internal/api"
	"hs-login/internal/repository"
	"hs-login/internal/service"
	"hs-login/pkg/server"
	"net/http"
)

func main() {
	fmt.Println("Starting server")
	db, err := server.InitPSQL()
	if err != nil {
		return
	}
	defer server.Close(db)

	// repositories
	authRepo := repository.NewAuthRepository(db)
	// services
	authService := service.NewAuthService(authRepo)
	//handler
	c := api.NewAuthHandler(authService)

	r := chi.NewRouter()
	r.Post("/login", c.Login)
	http.ListenAndServe(":3000", r)
}
