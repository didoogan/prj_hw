package main

import (
	"hw14/internal/api"
	store "hw14/internal/repositories"
	"hw14/internal/services"
	"log"
	"net/http"
)

func main() {
	userService := services.NewUserService(store.NewUserMemoryRepository())
	cacheService := services.NewCacheService(store.NewMemoryCacheRepository())
	tokenService := services.NewTokenService()

	appService := &services.AppService{
		User:  userService,
		Token: tokenService,
		Cache: cacheService,
		Auth: services.NewAuthService(
			userService, cacheService, tokenService,
		),
	}

	a := api.New(appService)

	r := a.CreateRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Printf("Server run error: %s", err)
	}
}
