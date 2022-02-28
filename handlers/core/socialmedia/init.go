package socialmedia

import (
	"github.com/riawaryati/mygram/backend/domain/general"
	"github.com/riawaryati/mygram/backend/usecase"
	"github.com/sirupsen/logrus"
)

type SocialMediaHandler struct {
	SocialMedia SocialMediaDataHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) SocialMediaHandler {
	return SocialMediaHandler{
		SocialMedia: newSocialMediaHandler(uc, conf, logger),
	}
}
