package user

import (
	"database/sql"
	"fmt"
	"strings"

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
		updated_at
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
		?, ?, ?, ?, NOW()
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
		username = ?`

	// uqFilterUsernameLowwer = `
	// 	lower(username) = ?`

	uqFilterPassword = `
		password = ?`

	// uqFilterAge = `
	// 	age = ?`

	uqFilterEmail = `
		email = ?`

	// uqFilterEmailLowwer = `
	// 	lower(email) = ?`
)

type UserDataRepoItf interface {
	GetByID(userID int) (*du.User, error)
	GetByUsername(username string) ([]*du.User, error)
	GetByEmailPassword(email string, password string) (*du.User, error)
	GetByEmail(email string) (*du.User, error)
	IsExistUser(email, password string) (bool, error)
	InsertUser(data du.CreateUser) (int, error)
	UpdateUser(data du.UpdateUser) error
	DeleteByID(userID int) error
}

func (ur UserDataRepo) GetByID(userID int) (*du.User, error) {
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

func (ur UserDataRepo) GetByEmailPassword(email string, password string) (*du.User, error) {
	var res du.User

	q := fmt.Sprintf("%s%s%s AND %s", uqSelectUser, uqWhere, uqFilterEmail, uqFilterPassword)
	query, args, err := ur.DBList.Backend.Read.In(q, strings.ToLower(email), password)
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

func (ur UserDataRepo) InsertUser(data du.CreateUser) (int, error) {
	param := make([]interface{}, 0)

	param = append(param, data.Username)
	param = append(param, data.Email)
	param = append(param, data.Password)
	param = append(param, data.Age)

	query, args, err := ur.DBList.Backend.Write.In(uqInsertUser, param...)
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

	var userID int
	err = res.Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (ur UserDataRepo) UpdateUser(data du.UpdateUser) error {
	var err error

	q := fmt.Sprintf("%s, %s, %s %s%s", uqUpdateUser, uqFilterEmail, uqFilterUsername, uqWhere, uqFilterUserID)

	query, args, err := ur.DBList.Backend.Read.In(q, strings.ToLower(data.Email), strings.ToLower(data.Username), data.ID)
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

func (ur UserDataRepo) DeleteByID(userID int) error {
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
