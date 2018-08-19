package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/tienducitt/go-restful/src/model"
	"github.com/tienducitt/go-restful/src/repository"
)

type UserHandler struct {
	UserRepo repository.IUserRepository
}

func NewUserHandler(userRepo repository.IUserRepository) *UserHandler {
	return &UserHandler{UserRepo: userRepo}
}

func (h *UserHandler) GetAll(req *http.Request) Response {
	users, err := h.UserRepo.GetAll()
	if err != nil {
		return Error(http.StatusInternalServerError, err)
	}

	return Success(users)
}

func (h *UserHandler) Create(req *http.Request) Response {
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		return Error(http.StatusBadRequest, err)
	}

	// validate
	err = h.UserRepo.Create(&user)
	if err != nil {
		return Error(http.StatusInternalServerError, err)
	}

	return Success(user)
}

func (h *UserHandler) Update(req *http.Request) Response {
	id := mux.Vars(req)["id"]
	userId, err := strconv.Atoi(id)

	if err != nil {
		return Error(http.StatusBadRequest, err)
	}

	var newUser model.User
	err = json.NewDecoder(req.Body).Decode(&newUser)

	if err != nil {
		return Error(http.StatusBadRequest, err)
	}

	user, err := h.UserRepo.Get(int64(userId))
	if err != nil {
		return Error(http.StatusInternalServerError, err)
	}
	if user == nil {
		return Error(http.StatusNotFound, fmt.Errorf("User not found"))
	}

	user.ID = int64(userId)

	err = h.UserRepo.Update(&newUser)
	if err != nil {
		return Error(http.StatusInternalServerError, err)
	}

	return Success(newUser)
}

func (h *UserHandler) Delete(req *http.Request) Response {
	id := mux.Vars(req)["id"]
	userId, err := strconv.Atoi(id)
	if err != nil {
		return Error(http.StatusBadRequest, err)
	}

	user, err := h.UserRepo.Get(int64(userId))
	if err != nil {
		return Error(http.StatusInternalServerError, err)
	}
	if user == nil {
		return Error(http.StatusNotFound, fmt.Errorf("User not found"))
	}

	err = h.UserRepo.Delete(user)
	if err != nil {
		return Error(http.StatusInternalServerError, err)
	}

	return Success(user)
}
