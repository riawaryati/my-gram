package core

import (
	"github.com/riawaryati/mygram/backend/domain/general"
	"github.com/riawaryati/mygram/backend/handlers/core/authorization"
	"github.com/riawaryati/mygram/backend/handlers/core/user"
	"github.com/riawaryati/mygram/backend/usecase"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Token  authorization.TokenHandler
	Public authorization.PublicHandler
	User   user.UserHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) Handler {
	return Handler{
		Token:  authorization.NewTokenHandler(conf, logger),
		Public: authorization.NewPublicHandler(conf, logger),
		User:   user.NewHandler(uc, conf, logger),
	}
}
