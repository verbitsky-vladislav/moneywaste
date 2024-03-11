package routes

import (
	"log"
	"moneywaste/package/handlers"
)

func RegisterUserRoutes() {
	userHandler := handlers.NewUserHandler()

	log.Printf("GROUP USER ROUTES")
	RegisterRoute("/user/update/", "PUT", userHandler.UpdateUser)
	RegisterRoute("/user/getOne/", "GET", userHandler.GetOneUser)
	RegisterRoute("/user/getAll/", "GET", userHandler.GetAllUsers)
	log.Printf("\n\n")
}
