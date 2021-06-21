package repository

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/entity/value"
)

type UserRepository interface {
	Get(ID int) (*entity.User, error)
	GetByLoginID(LoginID string) (*entity.User, error)
	GetByLoginIDAndPassword(LoginID string, Password value.Password) (*entity.User, error)
	Create(User entity.User) error
	Update(User entity.User) error
}
