package controllers

import (
	"crud-service/handlers"
	"crud-service/models"
	"crud-service/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		handlers.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	createUser, err := uc.userService.CreateUser(user)
	if err != nil {
		handlers.ErrorResponse(w, http.StatusInternalServerError, "Failed to create user")
		return
	}
	handlers.JSONResponse(w, http.StatusCreated, createUser)
}

func (uc *UserController) GetAllUser(w http.ResponseWriter, r *http.Request) {

	var user []models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	getAllUser, err := uc.userService.GetAllUser()
	if err != nil {
		handlers.ErrorResponse(w, http.StatusInternalServerError, "Internal Server Eror")
		return
	}
	handlers.JSONResponse(w, http.StatusCreated, getAllUser)
}
func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	getUserById, err := uc.userService.GetUserById(userID)
	if err != nil {
		handlers.ErrorResponse(w, http.StatusInternalServerError, "Internal Server Eror")
		return
	}
	handlers.JSONResponse(w, http.StatusCreated, getUserById)
}

func (uc *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	err := uc.userService.DeleteUserById(userID)
	if err != nil {
		handlers.ErrorResponse(w, http.StatusInternalServerError, "Internal Server Eror")
		return
	}
	handlers.JSONResponse(w, http.StatusCreated, map[string]string{"message": "User deleted"})
}
func (uc *UserController) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]

	err := uc.userService.UpdateUserById(userId)
	if err != nil {
		handlers.ErrorResponse(w, http.StatusInternalServerError, "Internal Server Eror")
		return
	}
	handlers.JSONResponse(w, http.StatusCreated, map[string]string{"message": "User updated"})
}
