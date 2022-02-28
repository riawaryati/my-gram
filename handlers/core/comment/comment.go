package comment

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	cg "github.com/riawaryati/mygram/backend/constants/general"
	du "github.com/riawaryati/mygram/backend/domain/comment"
	"github.com/riawaryati/mygram/backend/domain/general"
	"github.com/riawaryati/mygram/backend/handlers"
	"github.com/riawaryati/mygram/backend/usecase"
	uu "github.com/riawaryati/mygram/backend/usecase/comment"
	"github.com/sirupsen/logrus"
	"gopkg.in/dealancer/validate.v2"
)

type CommentDataHandler struct {
	Usecase uu.CommentDataUsecaseItf
	conf    *general.SectionService
	log     *logrus.Logger
}

func newCommentHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) CommentDataHandler {
	return CommentDataHandler{
		Usecase: uc.Comment.Comment,
		conf:    conf,
		log:     logger,
	}
}

func (ch CommentDataHandler) CreateComment(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	var param du.CommentRequest

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

	comment, err := ch.Usecase.CreateComment(param, accessToken)
	if err != nil {
		respData.Data = general.ResponseMessageData{Message: err.Error()}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status: cg.Success,
		Data:   comment,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
	return
}

func (ch CommentDataHandler) UpdateComment(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	message := ""
	commentidParam, ok := mux.Vars(req)["commentId"]

	if !ok {
		message = "Url Param 'commentId' is missing"
		respData.Data = &handlers.ResponseMessageData{
			Message: cg.Fail,
		}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	commentid, err := strconv.Atoi(commentidParam)
	if err != nil {
		message = "Invalid param comment id"

		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	var param du.UpdateCommentRequest

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

	comment, err := ch.Usecase.UpdateComment(param, commentid, accessToken)
	if err != nil {
		respData.Data = general.ResponseMessageData{Message: err.Error()}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status: cg.Success,
		Data:   comment,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
	return
}

func (ch CommentDataHandler) DeleteComment(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	message := ""
	commentIdParam, ok := mux.Vars(req)["commentId"]
	if !ok {
		message = "Url Param 'commentId' is missing"
		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	commentId, err := strconv.Atoi(commentIdParam)
	if err != nil {
		message = "Invalid param social media id"

		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	deleted, err := ch.Usecase.DeleteByID(commentId)

	if err != nil {
		message = err.Error()

		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	if !deleted {
		message = "Update comment gagal"

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

func (ch CommentDataHandler) GetComments(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	authorizationHeader := req.Header.Get("Authorization")
	accessToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	comments, err := ch.Usecase.GetCommentsByToken(accessToken)

	if err != nil {
		message := err.Error()

		respData.Data = general.ResponseMessageData{Message: message}
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status: cg.Success,
		Data:   comments,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
}
