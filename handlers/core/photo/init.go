package photo

import (
	"github.com/riawaryati/mygram/backend/domain/general"
	"github.com/riawaryati/mygram/backend/usecase"
	"github.com/sirupsen/logrus"
)

type PhotoHandler struct {
	Photo PhotoDataHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) PhotoHandler {
	return PhotoHandler{
		Photo: newPhotoHandler(uc, conf, logger),
	}
}
