package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/riawaryati/mygram/backend/handlers/core"
)

func getSocialMedia(routerJWT *mux.Router, handler core.Handler) {
	routerJWT.HandleFunc("/socialmedias", handler.SocialMedia.SocialMedia.CreateSocialMedia).Methods(http.MethodPost)
	routerJWT.HandleFunc("/socialmedias", handler.SocialMedia.SocialMedia.GetSocialMedias).Methods(http.MethodGet)

	routerJWT.HandleFunc("/socialmedias/{socialMediaId}", handler.SocialMedia.SocialMedia.UpdateSocialMedia).Methods(http.MethodPut)
	routerJWT.HandleFunc("/socialmedias/{socialMediaId}", handler.SocialMedia.SocialMedia.DeleteSocialMedia).Methods(http.MethodDelete)
}
