package comment

import (
	"errors"

	du "github.com/riawaryati/mygram/backend/domain/comment"
	"github.com/riawaryati/mygram/backend/domain/general"
	"github.com/riawaryati/mygram/backend/infra"
	"github.com/riawaryati/mygram/backend/repo"
	ru "github.com/riawaryati/mygram/backend/repo/comment"
	rp "github.com/riawaryati/mygram/backend/repo/photo"
	rr "github.com/riawaryati/mygram/backend/repo/user"
	"github.com/riawaryati/mygram/backend/utils"
	"github.com/sirupsen/logrus"
)

type CommentDataUsecaseItf interface {
	CreateComment(data du.CommentRequest, token string) (*du.CreateCommentResponse, error)
	UpdateComment(data du.UpdateCommentRequest, commentID int, token string) (*du.UpdateCommentResponse, error)
	GetCommentsByToken(token string) ([]du.CommentResponse, error)
	DeleteByID(commentId int) (bool, error)
}

type CommentDataUsecase struct {
	Repo      ru.CommentDataRepoItf
	RepoUser  rr.UserDataRepoItf
	RepoPhoto rp.PhotoDataRepoItf
	DBList    *infra.DatabaseList
	Conf      *general.SectionService
	Log       *logrus.Logger
}

func newCommentDataUsecase(r repo.Repo, conf *general.SectionService, logger *logrus.Logger, dbList *infra.DatabaseList) CommentDataUsecase {
	return CommentDataUsecase{
		Repo:      r.Comment.Comment,
		RepoUser:  r.User.User,
		RepoPhoto: r.Photo.Photo,
		Conf:      conf,
		Log:       logger,
		DBList:    dbList,
	}
}

func (uu CommentDataUsecase) CreateComment(data du.CommentRequest, token string) (*du.CreateCommentResponse, error) {
	// tx, err := uu.DBList.Backend.Write.Begin()
	// if err != nil {
	// 	return nil, err
	// }

	userID, err := utils.GetUserIDFromToken(token, uu.Conf.App.SecretKey)
	if err != nil {
		uu.Log.WithField("user id", userID).WithError(err).Error("fail to get user id from token")
		return nil, err
	}

	var reqData = du.CreateComment{
		UserID:  userID,
		PhotoID: data.PhotoID,
		Message: data.Message,
	}
	commentID, err := uu.Repo.InsertComment(reqData)
	if err != nil {
		return nil, errors.New("failed to insert comment")
	}

	comment, err := uu.Repo.GetByID(commentID)
	if err != nil {
		return nil, errors.New("failed to get comment")
	}

	commentResponse := du.CreateCommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
	}

	return &commentResponse, nil
}

func (uu CommentDataUsecase) GetCommentsByToken(token string) ([]du.CommentResponse, error) {

	userID, err := utils.GetUserIDFromToken(token, uu.Conf.App.SecretKey)
	if err != nil {
		uu.Log.WithField("user id", userID).WithError(err).Error("fail to get user id from token")
		return nil, err
	}

	comments, err := uu.Repo.GetListByUserID(userID)
	if err != nil {
		uu.Log.WithField("request", userID).WithError(err).Errorf("fail to checking is exist comment")
		return nil, err
	}

	if comments == nil {
		uu.Log.WithField("request", userID).Errorf("comment is not exist")
		return nil, errors.New("comment not exist")
	}

	var res []du.CommentResponse
	for _, comment := range comments {
		user, _ := uu.RepoUser.GetByID(comment.UserID)
		photo, _ := uu.RepoPhoto.GetByID(comment.PhotoID)

		var userComment du.UserComment
		var photoComment du.PhotoComment

		if user != nil {
			userComment.Email = user.Email
			userComment.ID = user.ID
			userComment.UserName = user.UserName
		}

		if photo != nil {
			photoComment.Caption = photo.Caption
			photoComment.ID = photo.ID
			photoComment.PhotoUrl = photo.PhotoUrl
			photoComment.Title = photo.Title
			photoComment.UserID = photo.UserID
		}

		commentRes := du.CommentResponse{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			UpdatedAt: *comment.UpdatedAt,
			CreatedAt: comment.CreatedAt,
			User:      userComment,
			Photo:     photoComment,
		}

		res = append(res, commentRes)
	}

	return res, nil
}

func (uu CommentDataUsecase) DeleteByID(commentID int) (bool, error) {

	err := uu.Repo.DeleteByID(commentID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (uu CommentDataUsecase) UpdateComment(data du.UpdateCommentRequest, commentID int, token string) (*du.UpdateCommentResponse, error) {
	// tx, err := uu.DBList.Backend.Write.Begin()
	// if err != nil {
	// 	return nil, err
	// }

	userID, err := utils.GetUserIDFromToken(token, uu.Conf.App.SecretKey)
	if err != nil {
		uu.Log.WithField("user id", userID).WithError(err).Error("fail to get user id from token")
		return nil, err
	}

	var reqData = du.UpdateComment{
		ID:      commentID,
		UserID:  userID,
		Message: data.Message,
	}

	err = uu.Repo.UpdateComment(reqData)
	if err != nil {
		return nil, err
	}

	comment, err := uu.Repo.GetByID(commentID)
	if err != nil {
		return nil, errors.New("failed to get comment")
	}

	commentRes := du.UpdateCommentResponse{
		ID:        comment.ID,
		UserID:    comment.UserID,
		UpdatedAt: *comment.UpdatedAt,
	}

	photo, _ := uu.RepoPhoto.GetByID(comment.PhotoID)

	if photo != nil {
		commentRes.Caption = photo.Caption
		commentRes.PhotoUrl = photo.PhotoUrl
		commentRes.Title = photo.Title
	}

	return &commentRes, nil
}
