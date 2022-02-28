package photo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	cg "github.com/riawaryati/mygram/backend/constants/general"
	"github.com/riawaryati/mygram/backend/domain/general"
	du "github.com/riawaryati/mygram/backend/domain/photo"
	"github.com/riawaryati/mygram/backend/handlers"
	"github.com/riawaryati/mygram/backend/usecase"
	uu "github.com/riawaryati/mygram/backend/usecase/photo"
	"github.com/sirupsen/logrus"
	"gopkg.in/dealancer/validate.v2"
)

type PhotoDataHandler struct {
	Usecase uu.PhotoDataUsecaseItf
	conf    *general.SectionService
	log     *logrus.Logger
}

func newPhotoHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) PhotoDataHandler {
	return PhotoDataHandler{
		Usecase: uc.Photo.Photo,
		conf:    conf,
		log:     logger,
	}
}

func (ch PhotoDataHandler) CreatePhoto(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	var param du.PhotoRequest

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respData.Data = general.ResponseMessageData{Message: cg.HandlerErrorRequestDataEmpty}
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &param)
	if err != nil {
		respData.Data = general.ResponseMessageData{Message: cg.HandlerErrorRequestDataNotValid}
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = validate.Validate(param)
	if err != nil {
		respData.Data = general.ResponseMessageData{Message: cg.HandlerErrorRequestDataFormatInvalid}
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	authorizationHeader := req.Header.Get("Authorization")
	accessToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	photo, err := ch.Usecase.CreatePhoto(param, accessToken)
	if err != nil {
		respData.Data = general.ResponseMessageData{Message: err.Error()}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status: cg.Success,
		Data:   photo,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
	return
}

func (ch PhotoDataHandler) UpdatePhoto(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	message := ""
	photoidParam, ok := mux.Vars(req)["photoId"]

	if !ok {
		message = "Url Param 'photoId' is missing"
		respData.Data = &handlers.ResponseMessageData{
			Message: cg.Fail,
		}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	photoid, err := strconv.ParseInt(photoidParam, 0, 64)
	if err != nil {
		message = "Invalid param photo id"

		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	var param du.PhotoRequest

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respData.Data = general.ResponseMessageData{Message: cg.HandlerErrorRequestDataEmpty}
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &param)
	if err != nil {
		respData.Data = general.ResponseMessageData{Message: cg.HandlerErrorRequestDataNotValid}
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = validate.Validate(param)
	if err != nil {
		respData.Data = general.ResponseMessageData{Message: cg.HandlerErrorRequestDataFormatInvalid}
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	authorizationHeader := req.Header.Get("Authorization")
	accessToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	photo, err := ch.Usecase.UpdatePhoto(param, photoid, accessToken)
	if err != nil {
		respData.Data = general.ResponseMessageData{Message: err.Error()}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status: cg.Success,
		Data:   photo,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
	return
}

func (ch PhotoDataHandler) DeletePhoto(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	message := ""
	photoIdParam, ok := mux.Vars(req)["photoId"]
	if !ok {
		message = "Url Param 'photoId' is missing"
		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	photoId, err := strconv.ParseInt(photoIdParam, 0, 64)
	if err != nil {
		message = "Invalid param social media id"

		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	deleted, err := ch.Usecase.DeleteByID(photoId)

	if err != nil {
		message = err.Error()

		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	if !deleted {
		message = "Update photo gagal"

		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status: cg.Success,
		Data:   message,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
}

func (ch PhotoDataHandler) GetPhotos(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	authorizationHeader := req.Header.Get("Authorization")
	accessToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	photos, err := ch.Usecase.GetPhotosByToken(accessToken)

	if err != nil {
		message := err.Error()

		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status: cg.Success,
		Data:   photos,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
}
