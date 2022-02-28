package socialmedia

import (
	"time"
)

type SocialMedia struct {
	ID             int        `json:"id" db:"id"`
	Name           string     `json:"name" db:"name"`
	SocialMediaUrl string     `json:"socialMediaUrl" db:"social_media_url"`
	UserID         int        `json:"userId" db:"user_id"`
	CreatedAt      time.Time  `json:"-" db:"created_at"`
	UpdatedAt      *time.Time `json:"-" db:"updated_at"`
}

type SocialMediaRequest struct {
	Name           string `json:"name" validate:"empty=false"`
	SocialMediaUrl string `json:"social_media_url" validate:"empty=false"`
}

type CreateSocialMedia struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         int    `json:"user_id"`
}

type UpdateSocialMedia struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         int    `json:"user_id"`
}

type CreateSocialMediaResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type UpdateSocialMediaResponse struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserID         int        `json:"user_id"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type SocialMediaGetResponse struct {
	SocialMedias []SocialMediaResponse `json:"social_medias"`
}
type SocialMediaResponse struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	SocialMediaUrl string          `json:"social_media_url"`
	UserID         int             `json:"user_id"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      *time.Time      `json:"updated_at"`
	User           UserSocialMedia `json:"User"`
}

type UserSocialMedia struct {
	ID              int    `json:"id"`
	UserName        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}
