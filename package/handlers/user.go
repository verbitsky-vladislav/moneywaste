package handlers

import "net/http"

type userHandler struct{}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *userHandler) GetOneUser(w http.ResponseWriter, r *http.Request) {

}

func (h *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

}
