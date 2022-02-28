package socialmedia

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	cg "github.com/riawaryati/mygram/backend/constants/general"
	"github.com/riawaryati/mygram/backend/domain/general"
	du "github.com/riawaryati/mygram/backend/domain/socialmedia"
	"github.com/riawaryati/mygram/backend/handlers"
	"github.com/riawaryati/mygram/backend/usecase"
	uu "github.com/riawaryati/mygram/backend/usecase/socialmedia"
	"github.com/sirupsen/logrus"
	"gopkg.in/dealancer/validate.v2"
)

type SocialMediaDataHandler struct {
	Usecase uu.SocialMediaDataUsecaseItf
	conf    *general.SectionService
	log     *logrus.Logger
}

func newSocialMediaHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) SocialMediaDataHandler {
	return SocialMediaDataHandler{
		Usecase: uc.SocialMedia.SocialMedia,
		conf:    conf,
		log:     logger,
	}
}

func (ch SocialMediaDataHandler) CreateSocialMedia(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	var param du.SocialMediaRequest

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

	socialmedia, err := ch.Usecase.CreateSocialMedia(param, accessToken)
	if err != nil {
		respData.Data = general.ResponseMessageData{Message: err.Error()}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status: cg.Success,
		Data:   socialmedia,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
	return
}

func (ch SocialMediaDataHandler) UpdateSocialMedia(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	message := ""
	socialmediaidParam, ok := mux.Vars(req)["socialmediaId"]

	if !ok {
		message = "Url Param 'socialmediaId' is missing"
		respData.Data = &handlers.ResponseMessageData{
			Message: cg.Fail,
		}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	socialmediaid, err := strconv.ParseInt(socialmediaidParam, 0, 64)
	if err != nil {
		message = "Invalid param socialmedia id"

		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	var param du.SocialMediaRequest

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

	socialmedia, err := ch.Usecase.UpdateSocialMedia(param, socialmediaid, accessToken)
	if err != nil {
		respData.Data = general.ResponseMessageData{Message: err.Error()}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status: cg.Success,
		Data:   socialmedia,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
	return
}

func (ch SocialMediaDataHandler) DeleteSocialMedia(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	message := ""
	socialmediaIdParam, ok := mux.Vars(req)["socialMediaId"]
	if !ok {
		message = "Url Param 'socialMediaId' is missing"
		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	socialmediaId, err := strconv.ParseInt(socialmediaIdParam, 0, 64)
	if err != nil {
		message = "Invalid param social media id"

		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	deleted, err := ch.Usecase.DeleteByID(socialmediaId)

	if err != nil {
		message = err.Error()

		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	if !deleted {
		message = "Update socialmedia gagal"

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

func (ch SocialMediaDataHandler) GetSocialMedias(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	authorizationHeader := req.Header.Get("Authorization")
	accessToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	socialMedias, err := ch.Usecase.GetSocialMediasByToken(accessToken)

	if err != nil {
		message := err.Error()

		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status: cg.Success,
		Data:   socialMedias,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
}
