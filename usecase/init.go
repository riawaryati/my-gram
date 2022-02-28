package usecase

import (
	"github.com/riawaryati/mygram/backend/domain/general"
	"github.com/riawaryati/mygram/backend/infra"
	"github.com/riawaryati/mygram/backend/repo"
	"github.com/riawaryati/mygram/backend/usecase/comment"
	"github.com/riawaryati/mygram/backend/usecase/photo"
	"github.com/riawaryati/mygram/backend/usecase/socialmedia"
	"github.com/riawaryati/mygram/backend/usecase/user"
	"github.com/sirupsen/logrus"
)

type Usecase struct {
	User        user.UserUsecase
	Photo       photo.PhotoUsecase
	SocialMedia socialmedia.SocialMediaUsecase
	Comment     comment.CommentUsecase
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) Usecase {
	return Usecase{
		User:        user.NewUsecase(repo, conf, dbList, logger),
		Photo:       photo.NewUsecase(repo, conf, dbList, logger),
		SocialMedia: socialmedia.NewUsecase(repo, conf, dbList, logger),
		Comment:     comment.NewUsecase(repo, conf, dbList, logger),
	}
}
