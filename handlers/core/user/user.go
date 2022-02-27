package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	cg "github.com/riawaryati/mygram/backend/constants/general"
	"github.com/riawaryati/mygram/backend/domain/general"
	du "github.com/riawaryati/mygram/backend/domain/user"
	"github.com/riawaryati/mygram/backend/handlers"
	"github.com/riawaryati/mygram/backend/usecase"
	uu "github.com/riawaryati/mygram/backend/usecase/user"
	"github.com/sirupsen/logrus"
	"gopkg.in/dealancer/validate.v2"
)

type UserDataHandler struct {
	Usecase uu.UserDataUsecaseItf
	conf    *general.SectionService
	log     *logrus.Logger
}

func newUserHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) UserDataHandler {
	return UserDataHandler{
		Usecase: uc.User.User,
		conf:    conf,
		log:     logger,
	}
}

func (ch UserDataHandler) RegisterUser(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	var param du.CreateUser

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

	user, err := ch.Usecase.RegisterUser(param)
	if err != nil {
		respData.Data = general.ResponseMessageData{Message: err.Error()}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status: cg.Success,
		Data:   user,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
	return
}

func (ch UserDataHandler) LoginUser(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	var param du.UserLoginRequest

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

	jwtToken, err := ch.Usecase.LoginUser(param)
	if err != nil {
		respData.Data = general.ResponseMessageData{Message: err.Error()}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status: cg.Success,
		Data:   jwtToken,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
	return
}

func (ch UserDataHandler) UpdateUser(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	message := ""
	useridParam, ok := mux.Vars(req)["userId"]

	if !ok {
		message = "Url Param 'userId' is missing"
		respData.Data = &handlers.ResponseMessageData{
			Message: cg.Fail,
		}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	userid, err := strconv.ParseInt(useridParam, 0, 64)
	if err != nil {
		message = "Invalid param user id"

		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	var param du.UpdateUserRequest

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

	userUpdate := du.UpdateUser{
		ID:       userid,
		Email:    param.Email,
		Username: param.Username,
	}

	user, err := ch.Usecase.UpdateUser(userUpdate)
	if err != nil {
		respData.Data = general.ResponseMessageData{Message: err.Error()}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status: cg.Success,
		Data:   user,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
	return
}

func (ch UserDataHandler) DeleteByID(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	message := ""
	userIdParam, ok := mux.Vars(req)["userId"]
	if !ok {
		message = "Url Param 'userId' is missing"
		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	userId, err := strconv.ParseInt(userIdParam, 0, 64)
	if err != nil {
		message = "Invalid param order id"

		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	deleted, err := ch.Usecase.DeleteByID(userId)

	if err != nil {
		message = err.Error()

		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	if !deleted {
		message = "Update user gagal"

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
