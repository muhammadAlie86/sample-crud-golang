package main

import (
	"crud-service/controllers"
	"crud-service/database"
	"crud-service/repositories"
	"crud-service/routes"
	"crud-service/services"
	"log"
	"net/http"
)

func main() {
	// Initialize database
	db := database.NewDatabase()
	db.InitDatabase()

	// Initialize repositories
	userRepository := repositories.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepository)

	// Initialize controllers
	userController := controllers.NewUserController(userService)

	// Set up router
	r := routes.SetupRouterUser(userController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Server started on port 8080")
	log.Fatal(server.ListenAndServe())
}
