package main

import (
	"fmt"
	"go-hm2/handlers"
	"go-hm2/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	userService := service.NewUserService()
	userHandler := handlers.NewUserHandler(userService)

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/users", userHandler.GetAll).Methods("GET")
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}", userHandler.GetByID).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	fmt.Println("Запускаем сервер на :8080")
	http.ListenAndServe(":8080", r)
}
