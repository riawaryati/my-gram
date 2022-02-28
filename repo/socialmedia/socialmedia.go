package socialmedia

import (
	"database/sql"
	"fmt"
	"time"

	du "github.com/riawaryati/mygram/backend/domain/socialmedia"
	"github.com/riawaryati/mygram/backend/infra"
)

type SocialMediaDataRepo struct {
	DBList *infra.DatabaseList
}

func newSocialMediaDataRepo(dbList *infra.DatabaseList) SocialMediaDataRepo {
	return SocialMediaDataRepo{
		DBList: dbList,
	}
}

const (
	uqSelectSocialMedia = `
	SELECT
		id,
		name,
		social_media_url,
		user_id,
		created_at,
		updated_at
	FROM
		social_medias`

	uqInsertSocialMedia = `
	INSERT INTO social_medias (
		name,
		social_media_url,
		user_id,
		created_at
	) VALUES (
		?, ?, ?, ?
	)
	RETURNING id`

	uqUpdateSocialMedia = `
	UPDATE 
		social_medias
	SET
		updated_at = NOW()`

	uqDeleteSocialMedia = `
		DELETE FROM 
			social_medias `

	uqWhere = `
	WHERE`

	uqFilterSocialMediaID = `
		id = ?`

	uqFilterName = `
		name = ?`

	uqFilterSocialMediaUrl = `
		social_media_url = ?`

	uqFilterUserID = `
		user_id = ?`
)

type SocialMediaDataRepoItf interface {
	GetByID(socialmediaID int) (*du.SocialMedia, error)
	GetListByUserID(userID int) ([]du.SocialMedia, error)
	InsertSocialMedia(data du.CreateSocialMedia) (int, error)
	UpdateSocialMedia(data du.UpdateSocialMedia) error
	DeleteByID(socialmediaID int) error
}

func (ur SocialMediaDataRepo) GetByID(socialmediaID int) (*du.SocialMedia, error) {
	var res du.SocialMedia

	q := fmt.Sprintf("%s%s%s", uqSelectSocialMedia, uqWhere, uqFilterSocialMediaID)
	query, args, err := ur.DBList.Backend.Read.In(q, socialmediaID)
	if err != nil {
		return nil, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res.ID == 0 {
		return nil, nil
	}

	return &res, nil
}

func (ur SocialMediaDataRepo) GetListByUserID(userID int) ([]du.SocialMedia, error) {
	var res []du.SocialMedia

	q := fmt.Sprintf("%s%s%s", uqSelectSocialMedia, uqWhere, uqFilterUserID)
	query, args, err := ur.DBList.Backend.Read.In(q, userID)
	if err != nil {
		return nil, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Select(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return res, nil
}

func (ur SocialMediaDataRepo) InsertSocialMedia(data du.CreateSocialMedia) (int, error) {
	param := make([]interface{}, 0)

	param = append(param, data.Name)
	param = append(param, data.SocialMediaUrl)
	param = append(param, data.UserID)

	param = append(param, time.Now().UTC())

	query, args, err := ur.DBList.Backend.Write.In(uqInsertSocialMedia, param...)
	if err != nil {
		return 0, err
	}

	query = ur.DBList.Backend.Write.Rebind(query)

	var res *sql.Row
	// if tx == nil {
	res = ur.DBList.Backend.Write.QueryRow(query, args...)
	// } else {
	// 	res = tx.QueryRow(query, args...)
	// }

	if err != nil {
		return 0, err
	}

	err = res.Err()
	if err != nil {
		return 0, err
	}

	var socialmediaID int
	err = res.Scan(&socialmediaID)
	if err != nil {
		return 0, err
	}

	return socialmediaID, nil
}

func (ur SocialMediaDataRepo) UpdateSocialMedia(data du.UpdateSocialMedia) error {
	q := fmt.Sprintf("%s, %s, %s %s%s", uqUpdateSocialMedia, uqFilterName, uqFilterSocialMediaUrl, uqWhere, uqFilterSocialMediaID)

	query, args, err := ur.DBList.Backend.Read.In(q, data.Name, data.SocialMediaUrl, data.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	query = ur.DBList.Backend.Write.Rebind(query)
	_, err = ur.DBList.Backend.Write.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (ur SocialMediaDataRepo) DeleteByID(socialmediaID int) error {
	var err error

	q := fmt.Sprintf("%s %s %s", uqDeleteSocialMedia, uqWhere, uqFilterSocialMediaID)

	query, args, err := ur.DBList.Backend.Read.In(q, socialmediaID)
	if err != nil {
		return err
	}

	query = ur.DBList.Backend.Write.Rebind(query)
	_, err = ur.DBList.Backend.Write.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
