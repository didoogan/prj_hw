package api

import (
	"encoding/json"
	"fmt"
	"hw14/internal/api/dto"
	"hw14/internal/entities"
	"net/http"
)

var defaultTtl = 60

func (a *API) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		a.makeResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	passwordOk, err := a.service.User.CheckPassword(&entities.UserWithPassword{
		Login:    loginRequest.Login,
		Password: loginRequest.Password,
	})

	if err != nil {
		a.makeResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !passwordOk {
		a.makeResponse(w, "wrong login or password", http.StatusBadRequest)
		return
	}

	token, exists, err := a.service.Cache.Get(loginRequest.Login)
	if err != nil {
		a.makeResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !exists {
		token, err = a.service.Token.Generate(&entities.User{
			Login: loginRequest.Login,
		})

		if err != nil {
			a.makeResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = a.service.Cache.Set(loginRequest.Login, token, defaultTtl)
	}

	loginResponse := dto.LoginResponse{Token: token}
	a.makeResponse(w, loginResponse, http.StatusOK)
}

func (a *API) TokenInfoHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	token := queryParams.Get("token")
	if token == "" {
		a.makeResponse(w, "query parameter `token` should be provided", http.StatusBadRequest)
		return
	}

	jwt, err := a.service.Token.VerifyToken(token)
	if err != nil {
		a.makeResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("jwt :", jwt)

	claims, err := a.service.Token.ExtractClaims(jwt)
	if err != nil {
		a.makeResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	a.makeResponse(w, claims, http.StatusOK)
}
