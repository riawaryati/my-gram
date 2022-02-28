package photo

import (
	"github.com/riawaryati/mygram/backend/domain/general"
	"github.com/riawaryati/mygram/backend/infra"
	"github.com/riawaryati/mygram/backend/repo"
	"github.com/sirupsen/logrus"
)

type PhotoUsecase struct {
	Photo PhotoDataUsecaseItf
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) PhotoUsecase {
	return PhotoUsecase{
		Photo: newPhotoDataUsecase(repo, conf, logger, dbList),
	}
}
