package photo

import (
	"database/sql"
	"fmt"
	"time"

	du "github.com/riawaryati/mygram/backend/domain/photo"
	"github.com/riawaryati/mygram/backend/infra"
)

type PhotoDataRepo struct {
	DBList *infra.DatabaseList
}

func newPhotoDataRepo(dbList *infra.DatabaseList) PhotoDataRepo {
	return PhotoDataRepo{
		DBList: dbList,
	}
}

const (
	uqSelectPhoto = `
	SELECT
		id,
		title,
		caption,
		photo_url,
		user_id,
		created_at,
		updated_at
	FROM
		photos`

	uqInsertPhoto = `
	INSERT INTO photos (
		title,
		caption,
		photo_url,
		user_id,
		created_at
	) VALUES (
		?, ?, ?, ?, ?
	)
	RETURNING id`

	uqUpdatePhoto = `
	UPDATE 
		photos
	SET
		updated_at = NOW()`

	uqDeletePhoto = `
		DELETE FROM 
			photos `

	uqSelectExist = `
		SELECT EXISTS`

	uqWhere = `
	WHERE`

	uqFilterPhotoID = `
		id = ?`

	uqFilterTitle = `
		title = ?`

	uqFilterPhotoUrl = `
		photo_url = ?`

	uqFilterUserID = `
		user_id = ?`

	uqFilterCaption = `
		caption = ?`
)

type PhotoDataRepoItf interface {
	GetByID(photoID int) (*du.Photo, error)
	GetListByUserID(userID int) ([]du.Photo, error)
	InsertPhoto(data du.CreatePhoto) (int, error)
	UpdatePhoto(data du.UpdatePhoto) error
	DeleteByID(photoID int) error
}

func (ur PhotoDataRepo) GetByID(photoID int) (*du.Photo, error) {
	var res du.Photo

	q := fmt.Sprintf("%s%s%s", uqSelectPhoto, uqWhere, uqFilterPhotoID)
	query, args, err := ur.DBList.Backend.Read.In(q, photoID)
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

func (ur PhotoDataRepo) GetListByUserID(userID int) ([]du.Photo, error) {
	//fmt.Println("iniiih ")
	var res []du.Photo

	q := fmt.Sprintf("%s%s%s", uqSelectPhoto, uqWhere, uqFilterUserID)
	query, args, err := ur.DBList.Backend.Read.In(q, userID)
	if err != nil {
		return nil, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	//fmt.Println("kesini ")
	err = ur.DBList.Backend.Read.Select(&res, query, args...)
	//fmt.Println("result ", res)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// if res == nil {
	// 	return nil, sql.ErrNoRows
	// }

	return res, nil
}

func (ur PhotoDataRepo) InsertPhoto(data du.CreatePhoto) (int, error) {
	param := make([]interface{}, 0)

	param = append(param, data.Title)
	param = append(param, data.Caption)
	param = append(param, data.PhotoUrl)
	param = append(param, data.UserID)

	param = append(param, time.Now().UTC())

	query, args, err := ur.DBList.Backend.Write.In(uqInsertPhoto, param...)
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

	var photoID int
	err = res.Scan(&photoID)
	if err != nil {
		return 0, err
	}

	return photoID, nil
}

func (ur PhotoDataRepo) UpdatePhoto(data du.UpdatePhoto) error {
	var err error

	q := fmt.Sprintf("%s, %s, %s, %s %s%s", uqUpdatePhoto, uqFilterCaption, uqFilterTitle, uqFilterPhotoUrl, uqWhere, uqFilterPhotoID)

	query, args, err := ur.DBList.Backend.Read.In(q, data.Caption, data.Title, data.PhotoUrl, data.ID)
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

func (ur PhotoDataRepo) DeleteByID(photoID int) error {
	var err error

	q := fmt.Sprintf("%s %s %s", uqDeletePhoto, uqWhere, uqFilterPhotoID)

	query, args, err := ur.DBList.Backend.Read.In(q, photoID)
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
