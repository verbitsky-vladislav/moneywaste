package handlers

import "net/http"

type authHandler struct{}

func NewAuthHandler() *authHandler {
	return &authHandler{}
}

func (h *authHandler) SignIn(w http.ResponseWriter, r *http.Request) {

}

func (h *authHandler) SignUp(w http.ResponseWriter, r *http.Request) {

}

func (h *authHandler) Logout(w http.ResponseWriter, r *http.Request) {

}
