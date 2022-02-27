package photo

import (
	"time"
)

type Comment struct {
	ID        int64      `json:"id" db:"id"`
	UserID    string     `json:"userId" db:"user_id"`
	PhotoID   string     `json:"photoId" db:"photo_id"`
	Message   string     `json:"message" db:"message"`
	CreatedAt time.Time  `json:"-" db:"created_at"`
	UpdatedAt *time.Time `json:"-" db:"updated_at"`
}
