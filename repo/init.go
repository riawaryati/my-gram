package repo

import (
	"github.com/riawaryati/mygram/backend/infra"
	"github.com/riawaryati/mygram/backend/repo/comment"
	"github.com/riawaryati/mygram/backend/repo/photo"
	"github.com/riawaryati/mygram/backend/repo/socialmedia"
	"github.com/riawaryati/mygram/backend/repo/user"
	"github.com/sirupsen/logrus"
)

type Repo struct {
	User        user.UserRepo
	Comment     comment.CommentRepo
	Photo       photo.PhotoRepo
	SocialMedia socialmedia.SocialMediaRepo
}

func NewRepo(db *infra.DatabaseList, logger *logrus.Logger) Repo {
	return Repo{
		User: user.NewMasterRepo(db, logger),
	}
}
