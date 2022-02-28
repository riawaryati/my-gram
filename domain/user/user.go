package user

import (
	"time"
)

type User struct {
	ID        int        `json:"id" db:"id"`
	UserName  string     `json:"username" db:"username"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"password" db:"password"`
	Age       int        `json:"age" db:"age"`
	CreatedAt time.Time  `json:"-" db:"created_at"`
	UpdatedAt *time.Time `json:"-" db:"updated_at"`
}

type CreateUser struct {
	Age      int    `json:"age" validate:"gt=8"`
	Email    string `json:"email" validate:"empty=false & format=email"`
	Password string `json:"password" validate:"empty=false"`
	Username string `json:"username" validate:"empty=false"`
}

type CreateUserResponse struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"empty=false & format=email"`
	Password string `json:"password" validate:"empty=false"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" validate:"empty=false & format=email"`
	Username string `json:"username" validate:"empty=false"`
}

type UpdateUser struct {
	ID       int    `json:"id"`
	Email    string `json:"email" validate:"empty=false & format=email"`
	Username string `json:"username" validate:"empty=false"`
}

type UpdateUserResponse struct {
	ID        int        `json:"id"`
	UserName  string     `json:"username"`
	Email     string     `json:"email"`
	Age       int        `json:"age"`
	UpdatedAt *time.Time `json:"updated_at"`
}
