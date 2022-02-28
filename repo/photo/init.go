package photo

import (
	"github.com/riawaryati/mygram/backend/infra"
	"github.com/sirupsen/logrus"
)

type PhotoRepo struct {
	Photo PhotoDataRepoItf
}

func NewMasterRepo(db *infra.DatabaseList, logger *logrus.Logger) PhotoRepo {
	return PhotoRepo{
		Photo: newPhotoDataRepo(db),
	}
}
