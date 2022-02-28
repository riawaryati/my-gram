package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/riawaryati/mygram/backend/handlers/core"
)

func getComment(routerJWT *mux.Router, handler core.Handler) {
	routerJWT.HandleFunc("/comments", handler.Comment.Comment.CreateComment).Methods(http.MethodPost)
	routerJWT.HandleFunc("/comments", handler.Comment.Comment.GetComments).Methods(http.MethodGet)

	routerJWT.HandleFunc("/comments/{commentId}", handler.Comment.Comment.UpdateComment).Methods(http.MethodPut)
	routerJWT.HandleFunc("/comments/{commentId}", handler.Comment.Comment.DeleteComment).Methods(http.MethodDelete)
}
