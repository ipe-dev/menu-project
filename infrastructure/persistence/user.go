package persistence

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/infrastructure/database"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

func (p userPersistence) Create(User entity.User) error {
	tx := database.Db.Begin()
	err := tx.Create(User).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}
func (p userPersistence) Update(User entity.User) error {
	tx := database.Db.Begin()
	err := tx.Updates(User).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}
func (p userPersistence) Get(ID int) entity.User {
	Db := database.Db
	var user entity.User
	Db.Model(user).Where("id = ?", ID).Find(&user)

	return user
}
func (p userPersistence) Login(LoginID string, Password string) (bool, entity.User) {
	Db := database.Db
	var user entity.User
	var isLogin bool
	err := Db.Model(user).Where("login_id = ?", LoginID).Where("password = ?", Password).Find(&user).Error
	if user.ID != 0 && err == nil {
		isLogin = true
	} else {
		isLogin = false
	}
	return isLogin, user
}