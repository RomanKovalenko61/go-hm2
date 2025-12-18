package main

import (
	"fmt"
	"go-hm2/handlers"
	"go-hm2/metrics"
	"go-hm2/service"
	"go-hm2/utils"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	userService := service.NewUserService()
	userHandler := handlers.NewUserHandler(userService)

	r := mux.NewRouter()

	r.Use(utils.RateLimitMiddleware)
	r.Use(metrics.Handler)

	r.Handle("/metrics", promhttp.Handler())

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/users", userHandler.GetAll).Methods("GET")
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	fmt.Println("Запускаем сервер на :8080")
	fmt.Println("Метрики доступны по /metrics")
	http.ListenAndServe(":8080", r)
}
