package repository

import "github.com/ipe-dev/menu_project/domain/entity"

type UserRepository interface {
	Get(ID int) entity.User
	Create(entity.User) error
	Update(entity.User) error
	Login(LoginID string, Password string) (bool, entity.User)
	Logout(ID int) bool
}
