package usecase

import (
	"github.com/riawaryati/mygram/backend/domain/general"
	"github.com/riawaryati/mygram/backend/infra"
	"github.com/riawaryati/mygram/backend/repo"
	"github.com/riawaryati/mygram/backend/usecase/user"
	"github.com/sirupsen/logrus"
)

type Usecase struct {
	User user.UserUsecase
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) Usecase {
	return Usecase{
		User: user.NewUsecase(repo, conf, dbList, logger),
	}
}
