package socialmedia

import (
	"errors"

	"github.com/riawaryati/mygram/backend/domain/general"
	du "github.com/riawaryati/mygram/backend/domain/socialmedia"
	"github.com/riawaryati/mygram/backend/infra"
	"github.com/riawaryati/mygram/backend/repo"
	ru "github.com/riawaryati/mygram/backend/repo/socialmedia"
	rr "github.com/riawaryati/mygram/backend/repo/user"
	"github.com/riawaryati/mygram/backend/utils"
	"github.com/sirupsen/logrus"
)

type SocialMediaDataUsecaseItf interface {
	CreateSocialMedia(data du.SocialMediaRequest, token string) (*du.CreateSocialMediaResponse, error)
	UpdateSocialMedia(data du.SocialMediaRequest, socialMediaId int, token string) (*du.UpdateSocialMediaResponse, error)
	GetSocialMediasByToken(token string) ([]du.SocialMediaResponse, error)
	DeleteByID(socialmediaId int) (bool, error)
}

type SocialMediaDataUsecase struct {
	Repo     ru.SocialMediaDataRepoItf
	RepoUser rr.UserDataRepoItf
	DBList   *infra.DatabaseList
	Conf     *general.SectionService
	Log      *logrus.Logger
}

func newSocialMediaDataUsecase(r repo.Repo, conf *general.SectionService, logger *logrus.Logger, dbList *infra.DatabaseList) SocialMediaDataUsecase {
	return SocialMediaDataUsecase{
		Repo:     r.SocialMedia.SocialMedia,
		RepoUser: r.User.User,
		Conf:     conf,
		Log:      logger,
		DBList:   dbList,
	}
}

func (uu SocialMediaDataUsecase) CreateSocialMedia(data du.SocialMediaRequest, token string) (*du.CreateSocialMediaResponse, error) {
	tx, err := uu.DBList.Backend.Write.Begin()
	if err != nil {
		return nil, err
	}

	userID, err := utils.GetUserIDFromToken(token, uu.Conf.App.SecretKey)
	if err != nil {
		uu.Log.WithField("user id", userID).WithError(err).Error("fail to get user id from token")
		return nil, err
	}

	var reqData = du.CreateSocialMedia{
		UserID:         userID,
		SocialMediaUrl: data.SocialMediaUrl,
		Name:           data.Name,
	}

	socialmediaID, err := uu.Repo.InsertSocialMedia(tx, reqData)
	if err != nil {
		return nil, errors.New("failed to insert socialmedia")
	}

	socialmedia, err := uu.Repo.GetByID(socialmediaID)
	if err != nil {
		return nil, errors.New("failed to get socialmedia")
	}

	socialmediaResponse := du.CreateSocialMediaResponse{
		ID:             socialmedia.ID,
		Name:           socialmedia.Name,
		SocialMediaUrl: socialmedia.SocialMediaUrl,
		UserID:         socialmedia.UserID,
		CreatedAt:      socialmedia.CreatedAt,
	}

	return &socialmediaResponse, nil
}

func (uu SocialMediaDataUsecase) GetSocialMediasByToken(token string) ([]du.SocialMediaResponse, error) {

	userID, err := utils.GetUserIDFromToken(token, uu.Conf.App.SecretKey)
	if err != nil {
		uu.Log.WithField("user id", userID).WithError(err).Error("fail to get user id from token")
		return nil, err
	}

	socialmedias, err := uu.Repo.GetListByUserID(userID)
	if err != nil {
		uu.Log.WithField("request", userID).WithError(err).Errorf("fail to checking is exist socialmedia")
		return nil, err
	}

	if socialmedias == nil {
		uu.Log.WithField("request", userID).Errorf("socialmedia is not exist")
		return nil, errors.New("socialmedia not exist")
	}

	var res []du.SocialMediaResponse
	for _, socialmedia := range socialmedias {
		user, _ := uu.RepoUser.GetByID(socialmedia.UserID)

		var userSocialMedia du.UserSocialMedia

		if user != nil {
			userSocialMedia.ProfileImageUrl = ""
			userSocialMedia.UserName = user.UserName
		}

		socialmediaRes := du.SocialMediaResponse{
			ID:        socialmedia.ID,
			Name:      socialmedia.Name,
			UserID:    socialmedia.UserID,
			UpdatedAt: socialmedia.UpdatedAt,
			CreatedAt: socialmedia.CreatedAt,
			User:      userSocialMedia,
		}

		res = append(res, socialmediaRes)
	}

	return res, nil
}

func (uu SocialMediaDataUsecase) DeleteByID(socialmediaID int) (bool, error) {

	err := uu.Repo.DeleteByID(socialmediaID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (uu SocialMediaDataUsecase) UpdateSocialMedia(data du.SocialMediaRequest, socialMediaID int, token string) (*du.UpdateSocialMediaResponse, error) {
	tx, err := uu.DBList.Backend.Write.Begin()
	if err != nil {
		return nil, err
	}

	userID, err := utils.GetUserIDFromToken(token, uu.Conf.App.SecretKey)
	if err != nil {
		uu.Log.WithField("user id", userID).WithError(err).Error("fail to get user id from token")
		return nil, err
	}

	var reqData = du.UpdateSocialMedia{
		ID:             socialMediaID,
		UserID:         userID,
		SocialMediaUrl: data.SocialMediaUrl,
		Name:           data.Name,
	}

	err = uu.Repo.UpdateSocialMedia(tx, reqData)
	if err != nil {
		return nil, err
	}

	socialmedia, err := uu.Repo.GetByID(socialMediaID)
	if err != nil {
		return nil, errors.New("failed to get socialmedia")
	}

	socialmediaRes := du.UpdateSocialMediaResponse{
		ID:             socialmedia.ID,
		Name:           socialmedia.Name,
		SocialMediaUrl: socialmedia.SocialMediaUrl,
		UserID:         socialmedia.UserID,
		UpdatedAt:      socialmedia.UpdatedAt,
	}

	return &socialmediaRes, nil
}
