package photo

import (
	"time"
)

type Photo struct {
	ID        int64      `json:"id" db:"id"`
	Title     string     `json:"title" db:"title"`
	Caption   string     `json:"caption" db:"caption"`
	PhotoUrl  int        `json:"photoUrl" db:"photo_url"`
	UserID    string     `json:"userId" db:"user_id"`
	CreatedAt time.Time  `json:"-" db:"created_at"`
	UpdatedAt *time.Time `json:"-" db:"updated_at"`
}
