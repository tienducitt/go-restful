package handler

import (
	"encoding/json"
	"net/http"

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
