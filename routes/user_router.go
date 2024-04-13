package routes

import (
	"crud-service/controllers"

	"github.com/gorilla/mux"
)

func SetupRouterUser(userControler *controllers.UserController) *mux.Router {

	r := mux.NewRouter()

	//r.Use(middlewares.AuthMiddleware)

	r.HandleFunc("/users", userControler.CreateUser).Methods("POST")
	r.HandleFunc("/users", userControler.GetAllUser).Methods("GET")
	r.HandleFunc("/users/{id}", userControler.GetUserById).Methods("GET")
	r.HandleFunc("/users/{id}", userControler.DeleteUserById).Methods("DELETE")
	r.HandleFunc("/users/{id}", userControler.DeleteUserById).Methods("PUT")

	return r
}
