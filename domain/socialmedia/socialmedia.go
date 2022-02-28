package socialmedia

import (
	"time"
)

type SocialMedia struct {
	ID             int64      `json:"id" db:"id"`
	Name           string     `json:"name" db:"name"`
	SocialMediaUrl string     `json:"socialMediaUrl" db:"social_media_url"`
	UserID         int64      `json:"userId" db:"user_id"`
	CreatedAt      time.Time  `json:"-" db:"created_at"`
	UpdatedAt      *time.Time `json:"-" db:"updated_at"`
}

type SocialMediaRequest struct {
	Name           string `json:"message"`
	SocialMediaUrl string `json:"social_media_url"`
}

type CreateSocialMedia struct {
	Name           string `json:"message"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         int64  `json:"user_id"`
}

type UpdateSocialMedia struct {
	ID             int64  `json:"id"`
	Name           string `json:"message"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         int64  `json:"user_id"`
}

type CreateSocialMediaResponse struct {
	ID             int64     `json:"id"`
	Name           string    `json:"message"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int64     `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type UpdateSocialMediaResponse struct {
	ID             int64      `json:"id"`
	Name           string     `json:"message"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserID         int64      `json:"user_id"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type SocialMediaResponse struct {
	ID             int64           `json:"id"`
	Name           string          `json:"message"`
	SocialMediaUrl string          `json:"social_media_url"`
	UserID         int64           `json:"UserId"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      *time.Time      `json:"updatedAt"`
	User           UserSocialMedia `json:"User"`
}

type UserSocialMedia struct {
	ID              int64  `json:"id"`
	UserName        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}
