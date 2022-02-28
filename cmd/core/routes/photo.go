package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/riawaryati/mygram/backend/handlers/core"
)

func getPhoto(routerJWT *mux.Router, handler core.Handler) {
	routerJWT.HandleFunc("/photos", handler.Photo.Photo.CreatePhoto).Methods(http.MethodPost)
	routerJWT.HandleFunc("/photos", handler.Photo.Photo.GetPhotos).Methods(http.MethodGet)

	routerJWT.HandleFunc("/photos/{photoId}", handler.Photo.Photo.UpdatePhoto).Methods(http.MethodPut)
	routerJWT.HandleFunc("/photos/{photoId}", handler.Photo.Photo.DeletePhoto).Methods(http.MethodDelete)
}
