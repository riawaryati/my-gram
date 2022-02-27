package photo

import (
	"time"
)

type Photo struct {
	ID             int64      `json:"id" db:"id"`
	Name           string     `json:"name" db:"name"`
	SocialMediaUrl string     `json:"socialMediaUrl" db:"social_media_url"`
	UserID         int        `json:"userId" db:"user_id"`
	CreatedAt      time.Time  `json:"-" db:"created_at"`
	UpdatedAt      *time.Time `json:"-" db:"updated_at"`
}
