package users

import (
	"bux-wallet/data/users"
	"time"
)

// User is a struct that contains user data.
type User struct {
	Email     string    `json:"email"`
	Xpriv     string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

// toUserDto converts User to UserDto.
func (user *User) toUserDto() *users.UserDto {
	return &users.UserDto{
		Email:     user.Email,
		Xpriv:     user.Xpriv,
		CreatedAt: user.CreatedAt,
	}
}
