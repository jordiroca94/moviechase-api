package user

import (
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router, handler *UserHandler) {
	router.HandleFunc("/login", handler.handleLogin).Methods("POST")
	router.HandleFunc("/register", handler.handleRegister).Methods("POST")
	router.HandleFunc("/user/{id}", handler.handleGetUser).Methods("GET")
	router.HandleFunc("/user/update/{id}", handler.handleUpdateUser).Methods("POST")
}
