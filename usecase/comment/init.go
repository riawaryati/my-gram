package comment

import (
	"github.com/riawaryati/mygram/backend/domain/general"
	"github.com/riawaryati/mygram/backend/infra"
	"github.com/riawaryati/mygram/backend/repo"
	"github.com/sirupsen/logrus"
)

type CommentUsecase struct {
	Comment CommentDataUsecaseItf
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) CommentUsecase {
	return CommentUsecase{
		Comment: newCommentDataUsecase(repo, conf, logger, dbList),
	}
}
