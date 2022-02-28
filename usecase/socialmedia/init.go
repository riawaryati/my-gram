package socialmedia

import (
	"github.com/riawaryati/mygram/backend/domain/general"
	"github.com/riawaryati/mygram/backend/infra"
	"github.com/riawaryati/mygram/backend/repo"
	"github.com/sirupsen/logrus"
)

type SocialMediaUsecase struct {
	SocialMedia SocialMediaDataUsecaseItf
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) SocialMediaUsecase {
	return SocialMediaUsecase{
		SocialMedia: newSocialMediaDataUsecase(repo, conf, logger, dbList),
	}
}
