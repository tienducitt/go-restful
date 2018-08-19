package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tienducitt/go-restful/src/handler"
	"github.com/tienducitt/go-restful/src/repository"
	"github.com/tienducitt/go-restful/src/router/middleware"
)

type Router struct {
	userRepo repository.IUserRepository
}

func NewRouter(userRepo repository.IUserRepository) Router {
	return Router{userRepo: userRepo}
}

func (r Router) Route() http.Handler {
	router := mux.NewRouter()

	userHandler := handler.NewUserHandler(r.userRepo)

	mwChain := middleware.ChainMiddleware(middleware.LoggingMw)

	router.HandleFunc("/users", mwChain(handler.MakeHandler(userHandler.GetAll))).Methods("GET")
	router.HandleFunc("/users", mwChain(handler.MakeHandler(userHandler.Create))).Methods("POST")
	router.HandleFunc("/users/{id}", mwChain(handler.MakeHandler(userHandler.Update))).Methods("PUT")
	router.HandleFunc("/users/{id}", mwChain(handler.MakeHandler(userHandler.Delete))).Methods("DELETE")

	return router
}
