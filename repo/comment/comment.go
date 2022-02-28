package comment

import (
	"database/sql"
	"fmt"
	"time"

	du "github.com/riawaryati/mygram/backend/domain/comment"
	"github.com/riawaryati/mygram/backend/infra"
)

type CommentDataRepo struct {
	DBList *infra.DatabaseList
}

func newCommentDataRepo(dbList *infra.DatabaseList) CommentDataRepo {
	return CommentDataRepo{
		DBList: dbList,
	}
}

const (
	uqSelectComment = `
	SELECT
		id,
		message,
		photo_id,
		user_id,
		created_at,
		updated_at
	FROM
		comments`

	uqInsertComment = `
	INSERT INTO comments (
		message,
		photo_id,
		user_id,
		created_at
	) VALUES (
		?, ?, ?, ?
	)
	RETURNING id`

	uqUpdateComment = `
	UPDATE 
		comments
	SET
		updated_at = NOW()`

	uqDeleteComment = `
		DELETE FROM 
			comments `

	uqWhere = `
	WHERE`

	uqFilterCommentID = `
		id = ?`

	uqFilterUserID = `
		user_id = ?`

	uqFilterMessage = `
		message = ?`
)

type CommentDataRepoItf interface {
	GetByID(commentID int) (*du.Comment, error)
	GetListByUserID(userID int) ([]du.Comment, error)
	InsertComment(data du.CreateComment) (int, error)
	UpdateComment(data du.UpdateComment) error
	DeleteByID(commentID int) error
}

func (ur CommentDataRepo) GetByID(commentID int) (*du.Comment, error) {
	var res du.Comment

	q := fmt.Sprintf("%s%s%s", uqSelectComment, uqWhere, uqFilterCommentID)
	query, args, err := ur.DBList.Backend.Read.In(q, commentID)
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

func (ur CommentDataRepo) GetListByUserID(userID int) ([]du.Comment, error) {
	var res []du.Comment

	q := fmt.Sprintf("%s%s%s", uqSelectComment, uqWhere, uqFilterUserID)
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

func (ur CommentDataRepo) InsertComment(data du.CreateComment) (int, error) {
	param := make([]interface{}, 0)

	param = append(param, data.Message)
	param = append(param, data.PhotoID)
	param = append(param, data.UserID)

	param = append(param, time.Now().UTC())

	query, args, err := ur.DBList.Backend.Write.In(uqInsertComment, param...)
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

	var commentID int
	err = res.Scan(&commentID)
	if err != nil {
		return 0, err
	}

	return commentID, nil
}

func (ur CommentDataRepo) UpdateComment(data du.UpdateComment) error {
	var err error

	q := fmt.Sprintf("%s, %s %s%s", uqUpdateComment, uqFilterMessage, uqWhere, uqFilterCommentID)

	query, args, err := ur.DBList.Backend.Read.In(q, data.Message, data.ID)
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

func (ur CommentDataRepo) DeleteByID(commentID int) error {
	var err error

	q := fmt.Sprintf("%s %s %s", uqDeleteComment, uqWhere, uqFilterCommentID)

	query, args, err := ur.DBList.Backend.Read.In(q, commentID)
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
