package user

import (
	"github.com/riawaryati/mygram/backend/domain/general"
	"github.com/riawaryati/mygram/backend/usecase"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	User UserDataHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) UserHandler {
	return UserHandler{
		User: newUserHandler(uc, conf, logger),
	}
}
