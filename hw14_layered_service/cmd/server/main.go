package main

import (
	"hw14/internal/api"
	store "hw14/internal/repositories"
	"hw14/internal/services"
	"log"
	"net/http"
)

func main() {
	appService := &services.AppService{
		User:  services.NewUserService(store.NewUserMemoryRepository()),
		Token: services.NewTokenService(),
		Cache: services.NewCacheService(store.NewMemoryCacheRepository()),
	}

	a := api.New(appService)

	r := a.CreateRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Printf("Server run error: %s", err)
	}
}
