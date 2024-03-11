package main

import (
	"log"
	"moneywaste/package/routes"
	"net/http"
)

func main() {
	routes.Register()

	log.Println("Сервер запущен и доступен на http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
