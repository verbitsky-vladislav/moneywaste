package main

import (
	"moneywaste/internal/http"
	"moneywaste/repository"
)

func main() {
	repository.InitDB("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	handlers := http.NewHandlers()
	handlers.RunRouters()
}
