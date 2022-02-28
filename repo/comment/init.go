package comment

import (
	"github.com/riawaryati/mygram/backend/infra"
	"github.com/sirupsen/logrus"
)

type CommentRepo struct {
	Comment CommentDataRepoItf
}

func NewMasterRepo(db *infra.DatabaseList, logger *logrus.Logger) CommentRepo {
	return CommentRepo{
		Comment: newCommentDataRepo(db),
	}
}
