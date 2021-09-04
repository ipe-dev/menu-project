package entity

import (
	"time"

	"github.com/ipe-dev/menu_project/domain/entity/value"
	"github.com/ipe-dev/menu_project/errors"
)

type UserOption func(*User) error
type User struct {
	ID        int
	Name      string
	LoginID   value.LoginID
	Password  value.Password
	CreatedAt time.Time
	UpdatedAt time.Time
}

func UserIDOption(ID int) UserOption {
	return func(u *User) error {
		if ID != 0 {
			u.ID = ID
			return nil
		}
	}
}
func UserNameOption(Name string) UserOption {
	return func(u *User) error {
		if Name != "" {
			u.Name = Name
			return nil
		}
	}
}
func LoginIDOption(LoginID string) UserOption {
	return func(u *User) error {
		var err error
		if LoginID != "" {
			u.LoginID, err = value.NewLoginID(LoginID)
			return err
		}
		return errors.NewValidateError(nil, "ログインIDが入力されていません")
	}
}
func PasswordOption(Password string) UserOption {
	return func(u *User) error {
		var err error
		if Password != "" {
			u.Password, err = value.NewPassword(Password)
			return err
		}
		return errors.NewValidateError(nil, "パスワードが入力されていません")
	}
}
func NewUser(opts ...UserOption) (*User, error) {
	user := new(User)
	var err error
	for _, opt := range opts {
		err = opt(user)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}
