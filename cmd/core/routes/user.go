package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/riawaryati/mygram/backend/handlers/core"
)

func getUser(router, routerJWT *mux.Router, handler core.Handler) {
	router.HandleFunc("/users/register", handler.User.User.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc("/users/login", handler.User.User.LoginUser).Methods(http.MethodPost)

	routerJWT.HandleFunc("/users/{userId}", handler.User.User.UpdateUser).Methods(http.MethodPut)
	routerJWT.HandleFunc("/users", handler.User.User.DeleteUser).Methods(http.MethodDelete)
}
