package photo

import (
	"errors"
	"fmt"

	"github.com/riawaryati/mygram/backend/domain/general"
	du "github.com/riawaryati/mygram/backend/domain/photo"
	"github.com/riawaryati/mygram/backend/infra"
	"github.com/riawaryati/mygram/backend/repo"
	ru "github.com/riawaryati/mygram/backend/repo/photo"
	rr "github.com/riawaryati/mygram/backend/repo/user"
	"github.com/riawaryati/mygram/backend/utils"
	"github.com/sirupsen/logrus"
)

type PhotoDataUsecaseItf interface {
	CreatePhoto(data du.PhotoRequest, token string) (*du.CreatePhotoResponse, error)
	UpdatePhoto(data du.PhotoRequest, photoID int, token string) (*du.UpdatePhotoResponse, error)
	GetPhotosByToken(token string) ([]du.PhotoResponse, error)
	DeleteByID(photoId int) (bool, error)
}

type PhotoDataUsecase struct {
	Repo     ru.PhotoDataRepoItf
	RepoUser rr.UserDataRepoItf
	DBList   *infra.DatabaseList
	Conf     *general.SectionService
	Log      *logrus.Logger
}

func newPhotoDataUsecase(r repo.Repo, conf *general.SectionService, logger *logrus.Logger, dbList *infra.DatabaseList) PhotoDataUsecase {
	return PhotoDataUsecase{
		Repo:     r.Photo.Photo,
		RepoUser: r.User.User,
		Conf:     conf,
		Log:      logger,
		DBList:   dbList,
	}
}

func (uu PhotoDataUsecase) CreatePhoto(data du.PhotoRequest, token string) (*du.CreatePhotoResponse, error) {
	// tx, err := uu.DBList.Backend.Write.Begin()
	// if err != nil {
	// 	return nil, err
	// }

	userID, err := utils.GetUserIDFromToken(token, uu.Conf.App.SecretKey)
	if err != nil {
		uu.Log.WithField("user id", userID).WithError(err).Error("fail to get user id from token")
		return nil, err
	}

	var reqData = du.CreatePhoto{
		UserID:   userID,
		PhotoUrl: data.PhotoUrl,
		Title:    data.Title,
		Caption:  data.Caption,
	}

	photoID, err := uu.Repo.InsertPhoto(reqData)
	if err != nil {
		return nil, errors.New("failed to insert photo")
	}

	photo, err := uu.Repo.GetByID(photoID)
	if err != nil {
		return nil, errors.New("failed to get photo")
	}

	photoResponse := du.CreatePhotoResponse{
		ID:        photo.ID,
		Caption:   photo.Caption,
		Title:     photo.Title,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
	}

	return &photoResponse, nil
}

func (uu PhotoDataUsecase) GetPhotosByToken(token string) ([]du.PhotoResponse, error) {

	userID, err := utils.GetUserIDFromToken(token, uu.Conf.App.SecretKey)
	if err != nil {
		uu.Log.WithField("user id", userID).WithError(err).Error("fail to get user id from token")
		return nil, err
	}

	fmt.Println(userID)
	photos, err := uu.Repo.GetListByUserID(userID)
	if err != nil {
		uu.Log.WithField("request", userID).WithError(err).Errorf("fail to checking is exist photo")
		return nil, err
	}

	fmt.Println(photos)
	if photos == nil {
		uu.Log.WithField("request", userID).Errorf("photo is not exist")
		return nil, errors.New("photo not exist")
	}

	var res []du.PhotoResponse
	for _, photo := range photos {
		user, _ := uu.RepoUser.GetByID(photo.UserID)

		var userPhoto du.UserPhoto

		if user != nil {
			userPhoto.Email = user.Email
			userPhoto.UserName = user.UserName
		}

		photoRes := du.PhotoResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			UserID:    photo.UserID,
			UpdatedAt: photo.UpdatedAt,
			CreatedAt: photo.CreatedAt,
			User:      userPhoto,
		}

		res = append(res, photoRes)
	}

	return res, nil
}

func (uu PhotoDataUsecase) DeleteByID(photoID int) (bool, error) {

	err := uu.Repo.DeleteByID(photoID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (uu PhotoDataUsecase) UpdatePhoto(data du.PhotoRequest, photoID int, token string) (*du.UpdatePhotoResponse, error) {
	// tx, err := uu.DBList.Backend.Write.Begin()
	// if err != nil {
	// 	return nil, err
	// }

	userID, err := utils.GetUserIDFromToken(token, uu.Conf.App.SecretKey)
	if err != nil {
		uu.Log.WithField("user id", userID).WithError(err).Error("fail to get user id from token")
		return nil, err
	}

	var reqData = du.UpdatePhoto{
		ID:       photoID,
		UserID:   userID,
		PhotoUrl: data.PhotoUrl,
		Caption:  data.Caption,
		Title:    data.Title,
	}

	err = uu.Repo.UpdatePhoto(reqData)
	if err != nil {
		return nil, err
	}

	photo, err := uu.Repo.GetByID(photoID)
	if err != nil {
		return nil, errors.New("failed to get photo")
	}

	photoRes := du.UpdatePhotoResponse{
		ID:        photo.ID,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		Title:     photo.Title,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}

	return &photoRes, nil
}
