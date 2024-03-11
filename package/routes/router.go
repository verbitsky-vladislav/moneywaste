package routes

import (
	"log"
	"moneywaste/common/middleware"
	"net/http"
)

func Register() {
	RegisterAuthRoutes()
	RegisterUserRoutes()
}

func RegisterRoute(pattern string, method string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, middleware.Method(handler, method))
	log.Printf("Path : [%s] | Method : [%s]\n", pattern, method)
}
