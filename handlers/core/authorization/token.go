package authorization

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	cg "github.com/riawaryati/mygram/backend/constants/general"
	dg "github.com/riawaryati/mygram/backend/domain/general"
	"github.com/riawaryati/mygram/backend/handlers"
	"github.com/riawaryati/mygram/backend/utils"
	"github.com/sirupsen/logrus"
)

type TokenHandler struct {
	log  *logrus.Logger
	Conf *dg.SectionService
}

func NewTokenHandler(conf *dg.SectionService, logger *logrus.Logger) TokenHandler {
	utils.InitJWTConfig(conf.Authorization.JWT)
	return TokenHandler{
		log:  logger,
		Conf: conf,
	}
}

func (th TokenHandler) JWTValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		respData := handlers.ResponseData{
			Status: cg.Fail,
		}

		//List of URL that bypass this JWTValidator middleware
		if req.URL.Path == "/api/v1/renew-token" {
			next.ServeHTTP(res, req)
			return
		}

		authorizationHeader := req.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			th.log.Error(fmt.Errorf("Invalid Token Format"))
			respData.Data = handlers.ResponseMessageData{Message: "Invalid Token Format"}
			handlers.WriteResponse(res, respData, http.StatusBadRequest)
			return
		}
		accessToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		claims, err := utils.CheckAccessToken(accessToken)
		if err != nil {
			respData.Data = handlers.ResponseMessageData{Message: "Token expired"}
			handlers.WriteResponse(res, respData, http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(req.Context(), "session", claims["session"])
		req = req.WithContext(ctx)

		next.ServeHTTP(res, req)
	})
}
