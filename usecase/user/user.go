package user

import (
	"errors"
	"fmt"

	"github.com/riawaryati/mygram/backend/domain/general"
	du "github.com/riawaryati/mygram/backend/domain/user"
	"github.com/riawaryati/mygram/backend/infra"
	"github.com/riawaryati/mygram/backend/repo"
	ru "github.com/riawaryati/mygram/backend/repo/user"
	"github.com/riawaryati/mygram/backend/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserDataUsecaseItf interface {
	RegisterUser(data du.CreateUser) (*du.CreateUserResponse, error)
	LoginUser(data du.UserLoginRequest) (*general.JWTAccess, error)
	UpdateUser(data du.UpdateUser) (bool, error)
	DeleteByAccessToken(accessToken string) (bool, error)
}

type UserDataUsecase struct {
	Repo   ru.UserDataRepoItf
	DBList *infra.DatabaseList
	Conf   *general.SectionService
	Log    *logrus.Logger
}

func newUserDataUsecase(r repo.Repo, conf *general.SectionService, logger *logrus.Logger, dbList *infra.DatabaseList) UserDataUsecase {
	return UserDataUsecase{
		Repo:   r.User.User,
		Conf:   conf,
		Log:    logger,
		DBList: dbList,
	}
}

func (uu UserDataUsecase) RegisterUser(data du.CreateUser) (*du.CreateUserResponse, error) {
	tx, err := uu.DBList.Backend.Write.Begin()
	if err != nil {
		return nil, err
	}

	userID, err := uu.Repo.InsertUser(tx, data)
	if err != nil {
		return nil, errors.New("failed to insert user")
	}

	user, err := uu.Repo.GetByID(userID)
	if err != nil {
		return nil, errors.New("failed to get user")
	}

	userResponse := du.CreateUserResponse{
		Age:      user.Age,
		Email:    user.Email,
		ID:       user.ID,
		Username: user.UserName,
	}

	return &userResponse, nil
}

func (uu UserDataUsecase) LoginUser(data du.UserLoginRequest) (*general.JWTAccess, error) {
	passwordHash, err := hashPassword(data.Password)

	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to hash password")
		return nil, err
	}

	user, err := uu.Repo.GetByEmailPassword(data.Email, passwordHash)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to checking is exist user")
		return nil, err
	}

	if user == nil {
		uu.Log.WithField("request", utils.StructToString(data)).Errorf("user is not exist")
		return nil, errors.New("user not exist")
	}

	session, err := utils.GetEncrypt([]byte(uu.Conf.App.SecretKey), fmt.Sprintf("%v", user.ID))
	if err != nil {
		uu.Log.WithField("user id", user.ID).WithError(err).Error("fail to get token data from infra")
		return nil, err
	}

	accessToken, _, err := utils.GenerateJWT(session)
	if err != nil {
		uu.Log.WithField("user id", user.ID).WithError(err).Error("fail to get token data from infra")
		return nil, err
	}

	jwtToken := &general.JWTAccess{Token: accessToken}

	return jwtToken, nil
}

func (uu UserDataUsecase) DeleteByAccessToken(accessToken string) (bool, error) {

	userID, err := utils.GetUserIDFromToken(accessToken, uu.Conf.App.SecretKey)
	if err != nil {
		uu.Log.WithField("user id", userID).WithError(err).Error("fail to get user id from token")
		return false, err
	}

	err = uu.Repo.DeleteByID(userID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (uu UserDataUsecase) UpdateUser(data du.UpdateUser) (bool, error) {
	tx, err := uu.DBList.Backend.Write.Begin()
	if err != nil {
		return false, err
	}

	err = uu.Repo.UpdateUser(tx, data)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	return true, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
