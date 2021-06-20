package entity

import (
	"log"
	"time"

	"github.com/ipe-dev/menu_project/domain/entity/value"
	"github.com/ipe-dev/menu_project/domain/repository"
)

type UserOption func(*User)
type User struct {
	ID        int
	Name      string
	LoginID   string
	Password  value.Password
	CreatedAt time.Time
	UpdatedAt time.Time
}

func IDOption(ID int) UserOption {
	return func(u *User) {
		if ID != 0 {
			u.ID = ID
		}
	}
}
func UserNameOption(Name string) UserOption {
	return func(u *User) {
		if Name != "" {
			u.Name = Name
		}
	}
}
func LoginIDOption(LoginID string) UserOption {
	return func(u *User) {
		if LoginID != "" {
			u.LoginID = LoginID
		}
	}
}
func PasswordOption(Password string) UserOption {
	return func(u *User) {
		if Password != "" {
			u.Password = value.NewPassword(Password)
		}
	}
}
func NewUser(opts ...UserOption) *User {

	user := new(User)
	for _, opt := range opts {
		opt(user)
	}
	return user
}

func (u *User) LoginAuthentication() (*User, error) {
	GetUser, err := repository.UserRepository.GetByLoginIDAndPassword(u.LoginID, u.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return GetUser, err
}
