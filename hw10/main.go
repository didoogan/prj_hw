package main

import (
	"encoding/json"
	"hw10/entities"
	"hw10/store"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const apiPrefix = "api"
const apiVersion = "v1"

type TaskStorage interface {
	List() []*entities.Task
	Create(*entities.Task)
	Update(*entities.Task) error
	Delete(id int64) error
}

type Server struct {
	storage TaskStorage
}

func main() {

	srv := &Server{storage: store.NewMemStore()}
	r := mux.NewRouter()

	s := r.PathPrefix("/api/v1/tasks").Subrouter()

	s.HandleFunc("/", srv.listHandler).Methods(http.MethodGet)
	s.HandleFunc("/", srv.createHandler).Methods(http.MethodPost)
	s.HandleFunc("/{id}", srv.updateHandler).Methods(http.MethodPut)
	s.HandleFunc("/{id}", srv.deleteHandler).Methods(http.MethodDelete)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Printf("Server run error: %s", err)
	}
}

func writeResponse(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("error: %s", err)
	}
}

func (s *Server) listHandler(w http.ResponseWriter, r *http.Request) {
	tasks := s.storage.List()

	writeResponse(w, tasks, http.StatusOK)
}

func (s *Server) createHandler(w http.ResponseWriter, r *http.Request) {
	var task entities.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.storage.Create(&task)

	writeResponse(w, task, http.StatusCreated)
}

func (s *Server) updateHandler(w http.ResponseWriter, r *http.Request) {
	var task entities.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil || id != task.ID {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.storage.Update(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeResponse(w, task, http.StatusOK)
}

func (s *Server) deleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.storage.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
