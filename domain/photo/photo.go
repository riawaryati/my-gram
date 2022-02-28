package photo

import (
	"time"
)

type Photo struct {
	ID        int64      `json:"id" db:"id"`
	Title     string     `json:"title" db:"title"`
	Caption   string     `json:"caption" db:"caption"`
	PhotoUrl  string     `json:"photoUrl" db:"photo_url"`
	UserID    int64      `json:"userId" db:"user_id"`
	CreatedAt time.Time  `json:"-" db:"created_at"`
	UpdatedAt *time.Time `json:"-" db:"updated_at"`
}

type PhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type CreatePhoto struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   int64  `json:"user_id"`
}

type UpdatePhoto struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   int64  `json:"user_id"`
}

type CreatePhotoResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoResponse struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    int64      `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      UserPhoto  `json:"User"`
}

type UserPhoto struct {
	Email    string `json:"email"`
	UserName string `json:"username"`
}

type UpdatePhotoResponse struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    int64      `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}
