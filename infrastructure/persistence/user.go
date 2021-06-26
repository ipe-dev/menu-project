package persistence

import (
	"log"

	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/entity/value"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/infrastructure/database"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

func (p userPersistence) Create(User entity.User) error {
	tx := database.Db.Begin()
	err := tx.Table("users").Create(&User).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}
func (p userPersistence) Update(User entity.User) error {
	tx := database.Db.Begin()
	err := tx.Where("id = ?", User.ID).Updates(&User).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}
func (p userPersistence) Get(ID int) (*entity.User, error) {
	Db := database.Db
	var user entity.User
	err := Db.Model(user).Where("id = ?", ID).Find(&user).Error

	return &user, err
}
func (p userPersistence) GetByLoginID(LoginID string) (*entity.User, error) {
	Db := database.Db
	var user entity.User
	err := Db.Model(user).Where("login_id = ?", LoginID).Find(&user).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, err
}
func (p userPersistence) GetByLoginIDAndPassword(LoginID string, Password value.Password) (*entity.User, error) {
	Db := database.Db
	user := new(entity.User)
	err := Db.Model(user).Where("login_id = ?", LoginID).Where("password = ?", Password).Find(user).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, err
}
