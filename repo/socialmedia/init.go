package socialmedia

import (
	"github.com/riawaryati/mygram/backend/infra"
	"github.com/sirupsen/logrus"
)

type SocialMediaRepo struct {
	SocialMedia SocialMediaDataRepoItf
}

func NewMasterRepo(db *infra.DatabaseList, logger *logrus.Logger) SocialMediaRepo {
	return SocialMediaRepo{
		SocialMedia: newSocialMediaDataRepo(db),
	}
}
