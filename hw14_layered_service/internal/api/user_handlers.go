package api

import (
	"encoding/json"
	"hw14/internal/api/dto"
	"hw14/internal/entities"
	"net/http"
)

func (a *API) UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	var createUserRequest dto.CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&createUserRequest)
	if err != nil {
		a.makeResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	a.service.User.Save(&entities.UserWithPassword{
		Login:    createUserRequest.Login,
		Password: createUserRequest.Password,
	})

	a.makeResponse(w, "ok", http.StatusCreated)
}

func (a *API) UserListHandler(w http.ResponseWriter, r *http.Request) {
	users, err := a.service.User.List()

	if err != nil {
		a.makeResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a.makeResponse(w, users, http.StatusOK)
}
