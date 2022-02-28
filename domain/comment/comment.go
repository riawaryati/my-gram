package photo

import (
	"time"
)

type Comment struct {
	ID        int64      `json:"id" db:"id"`
	UserID    int64      `json:"userId" db:"user_id"`
	PhotoID   int64      `json:"photoId" db:"photo_id"`
	Message   string     `json:"message" db:"message"`
	CreatedAt time.Time  `json:"-" db:"created_at"`
	UpdatedAt *time.Time `json:"-" db:"updated_at"`
}

type CommentRequest struct {
	Message string `json:"message"`
	PhotoID int64  `json:"photo_id"`
}

type UpdateCommentRequest struct {
	Message string `json:"message"`
}

type CreateComment struct {
	Message string `json:"message"`
	PhotoID int64  `json:"photo_id"`
	UserID  int64  `json:"user_id"`
}

type UpdateComment struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
	// PhotoID int64  `json:"photo_id"`
	UserID int64 `json:"user_id"`
}

type CreateCommentResponse struct {
	ID        int64     `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int64     `json:"photo_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentResponse struct {
	ID        int64        `json:"id"`
	Message   string       `json:"message"`
	PhotoID   int64        `json:"photo_id"`
	UserID    int64        `json:"user_id"`
	UpdatedAt time.Time    `json:"updated_at"`
	CreatedAt time.Time    `json:"created_at"`
	User      UserComment  `json:"User"`
	Photo     PhotoComment `json:"Photo"`
}

type UserComment struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}

type PhotoComment struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   int64  `json:"user_id"`
}

// type UpdateCommentResponse struct {
// 	ID        int64     `json:"id"`
// 	Message   string    `json:"message"`
// 	PhotoID   string    `json:"photo_id"`
// 	UserID    string    `json:"user_id"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }

type UpdateCommentResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int64     `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
