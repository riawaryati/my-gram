package comment

import (
	"github.com/riawaryati/mygram/backend/domain/general"
	"github.com/riawaryati/mygram/backend/usecase"
	"github.com/sirupsen/logrus"
)

type CommentHandler struct {
	Comment CommentDataHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) CommentHandler {
	return CommentHandler{
		Comment: newCommentHandler(uc, conf, logger),
	}
}
