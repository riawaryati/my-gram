package user

import (
	"github.com/riawaryati/mygram/backend/domain/general"
	"github.com/riawaryati/mygram/backend/infra"
	"github.com/riawaryati/mygram/backend/repo"
	"github.com/sirupsen/logrus"
)

type UserUsecase struct {
	User UserDataUsecaseItf
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) UserUsecase {
	return UserUsecase{
		User: newUserDataUsecase(repo, conf, logger, dbList),
	}
}
