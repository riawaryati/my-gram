package photo

import (
	"time"
)

type Comment struct {
	ID        int        `json:"id" db:"id"`
	UserID    int        `json:"userId" db:"user_id"`
	PhotoID   int        `json:"photoId" db:"photo_id"`
	Message   string     `json:"message" db:"message"`
	CreatedAt time.Time  `json:"-" db:"created_at"`
	UpdatedAt *time.Time `json:"-" db:"updated_at"`
}

type CommentRequest struct {
	Message string `json:"message" validate:"empty=false"`
	PhotoID int    `json:"photo_id"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" validate:"empty=false"`
}

type CreateComment struct {
	Message string `json:"message"`
	PhotoID int    `json:"photo_id"`
	UserID  int    `json:"user_id"`
}

type UpdateComment struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	// PhotoID int  `json:"photo_id"`
	UserID int `json:"user_id"`
}

type CreateCommentResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentResponse struct {
	ID        int          `json:"id"`
	Message   string       `json:"message"`
	PhotoID   int          `json:"photo_id"`
	UserID    int          `json:"user_id"`
	UpdatedAt time.Time    `json:"updated_at"`
	CreatedAt time.Time    `json:"created_at"`
	User      UserComment  `json:"User"`
	Photo     PhotoComment `json:"Photo"`
}

type UserComment struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}

type PhotoComment struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}

// type UpdateCommentResponse struct {
// 	ID        int     `json:"id"`
// 	Message   string    `json:"message"`
// 	PhotoID   string    `json:"photo_id"`
// 	UserID    string    `json:"user_id"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }

type UpdateCommentResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
