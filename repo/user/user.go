package user

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	du "github.com/riawaryati/mygram/backend/domain/user"
	"github.com/riawaryati/mygram/backend/infra"
)

type UserDataRepo struct {
	DBList *infra.DatabaseList
}

func newUserDataRepo(dbList *infra.DatabaseList) UserDataRepo {
	return UserDataRepo{
		DBList: dbList,
	}
}

const (
	uqSelectUser = `
	SELECT
		id,
		username,
		email,
		password,
		age,
		created_at,
		updated_at,
		updated_by
	FROM
		users`

	uqInsertUser = `
	INSERT INTO users (
		username,
		email,
		password,
		age,
		created_at
	) VALUES (
		?, ?, ?, ?, ?
	)
	RETURNING id`

	uqUpdateUser = `
	UPDATE 
		users
	SET
		updated_at = NOW()`

	uqDeleteUser = `
		DELETE FROM 
			users `

	uqSelectExist = `
		SELECT EXISTS`

	uqWhere = `
	WHERE`

	uqFilterUserID = `
		id = ?`

	uqFilterUsername = `
		lower(username) = ?`

	uqFilterPassword = `
		password = ?`

	// uqFilterAge = `
	// 	age = ?`

	uqFilterEmail = `
		lower(email) = ?`
)

type UserDataRepoItf interface {
	GetByID(userID int64) (*du.User, error)
	GetByUsername(username string) ([]*du.User, error)
	GetByEmail(email string) (*du.User, error)
	IsExistUser(email, password string) (bool, error)
	InsertUser(tx *sql.Tx, data du.CreateUser) (int64, error)
	UpdateUser(tx *sql.Tx, data du.UpdateUser) error
	DeleteByID(userID int64) error
}

func (ur UserDataRepo) GetByID(userID int64) (*du.User, error) {
	var res du.User

	q := fmt.Sprintf("%s%s%s", uqSelectUser, uqWhere, uqFilterUserID)
	query, args, err := ur.DBList.Backend.Read.In(q, userID)
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

func (ur UserDataRepo) GetByUsername(username string) ([]*du.User, error) {
	var res []*du.User

	q := fmt.Sprintf("%s%s%s", uqSelectUser, uqWhere, uqFilterUsername)
	query, args, err := ur.DBList.Backend.Read.In(q, strings.ToLower(username))
	if err != nil {
		return nil, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Select(&res, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ur UserDataRepo) GetByEmail(email string) (*du.User, error) {
	var res du.User

	q := fmt.Sprintf("%s%s%s", uqSelectUser, uqWhere, uqFilterEmail)
	query, args, err := ur.DBList.Backend.Read.In(q, strings.ToLower(email))
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

func (ur UserDataRepo) IsExistUser(email, password string) (bool, error) {
	var isExist bool

	q := fmt.Sprintf("%s(%s%s%s AND %s)", uqSelectExist, uqSelectUser, uqWhere, uqFilterEmail, uqFilterPassword)

	query, args, err := ur.DBList.Backend.Read.In(q, strings.ToLower(email), password)
	if err != nil {
		return isExist, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Get(&isExist, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return isExist, err
	}

	return isExist, nil
}

func (ur UserDataRepo) InsertUser(tx *sql.Tx, data du.CreateUser) (int64, error) {
	param := make([]interface{}, 0)

	param = append(param, data.Username)
	param = append(param, data.Email)
	param = append(param, data.Password)
	param = append(param, data.Age)

	param = append(param, time.Now().UTC())

	query, args, err := ur.DBList.Backend.Write.In(uqInsertUser, param...)
	if err != nil {
		return 0, err
	}

	query = ur.DBList.Backend.Write.Rebind(query)

	var res *sql.Row
	if tx == nil {
		res = ur.DBList.Backend.Write.QueryRow(query, args...)
	} else {
		res = tx.QueryRow(query, args...)
	}

	if err != nil {
		return 0, err
	}

	err = res.Err()
	if err != nil {
		return 0, err
	}

	var userID int64
	err = res.Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (ur UserDataRepo) UpdateUser(tx *sql.Tx, data du.UpdateUser) error {
	var err error

	q := fmt.Sprintf("%s, %s, %s %s%s", uqUpdateUser, uqFilterEmail, uqFilterUsername, uqWhere, uqFilterUserID)

	query, args, err := ur.DBList.Backend.Read.In(q, strings.ToLower(data.Username), strings.ToLower(data.Username), data.ID)
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

func (ur UserDataRepo) DeleteByID(userID int64) error {
	var err error

	q := fmt.Sprintf("%s %s %s", uqDeleteUser, uqWhere, uqFilterUserID)

	query, args, err := ur.DBList.Backend.Read.In(q, userID)
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
