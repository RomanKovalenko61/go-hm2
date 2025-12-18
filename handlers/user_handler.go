package handlers

import (
	"encoding/json"
	"go-hm2/models"
	"go-hm2/service"
	"go-hm2/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		go utils.AuditUserFailedAction("CREATE FAILED!", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	savedUser := handler.userService.Create(&req)

	go utils.AuditUserAction("CREATE USER", savedUser.ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(savedUser)
}

func (handler *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		go utils.AuditUserFailedAction("GET USER FAILED!", "ID is not a number")
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := handler.userService.GetById(id)
	if err != nil {
		go utils.AuditUserFailedAction("GET USER FAILED!", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	go utils.AuditUserAction("GET USER", user.ID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (handler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		go utils.AuditUserFailedAction("UPDATE FAILED!", err.Error())
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var req models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		go utils.AuditUserFailedAction("UPDATE FAILED!", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedUser, err := handler.userService.Update(id, &req)
	if err != nil {
		go utils.AuditUserFailedAction("UPDATE FAILED!", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	go utils.AuditUserAction("UPDATE USER", updatedUser.ID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		go utils.AuditUserFailedAction("DELETE FAILED!", err.Error())
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := handler.userService.Delete(id); err != nil {
		go utils.AuditUserFailedAction("DELETE FAILED!", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	go utils.AuditUserAction("DELETE USER", id)
	w.WriteHeader(http.StatusNoContent)
}

func (handler *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users := handler.userService.GetAll()

	go utils.AuditUserAction("GET ALL USERS", 0)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
