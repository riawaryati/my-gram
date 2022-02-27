package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/riawaryati/mygram/backend/domain/general"
	"github.com/riawaryati/mygram/backend/handlers/core"
)

func getUser(router, routerJWT *mux.Router, conf *general.SectionService, handler core.Handler) {
	router.HandleFunc("/users/register", handler.User.User.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc("/users/login", handler.User.User.LoginUser).Methods(http.MethodPost)

	router.HandleFunc("/users/{userId}", handler.User.User.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{userid}", handler.User.User.DeleteByID).Methods(http.MethodDelete)
}
