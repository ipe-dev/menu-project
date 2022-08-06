package persistence

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/errors"
	"github.com/ipe-dev/menu_project/infrastructure/database"
)

type menuPersistence struct{}

func NewMenuPersistence() repository.MenuRepository {
	return &menuPersistence{}
}
func (p menuPersistence) BulkCreate(menus []entity.Menu) error {
	tx := database.Db.Begin()
	err := tx.Model(&menus).Create(menus).Error
	if err != nil {
		tx.Rollback()
		return errors.NewInfraError(err, menus)
	}
	tx.Commit()
	return nil
}
func (p menuPersistence) BulkUpdate(menus []entity.Menu) error {
	tx := database.Db.Begin()
	var err error
	for _, menu := range menus {
		err = tx.Model(&menu).Updates(menu).Error
		if err != nil {
			tx.Rollback()
			return errors.NewInfraError(err, menu)
		}
	}
	tx.Commit()
	return nil
}
func (p menuPersistence) GetByID(id int) (entity.Menu, error) {
	var menu entity.Menu
	Db := database.Db
	err := Db.Table("menus").Where("id = ?", id).First(&menu).Error
	if err != nil {
		return menu, errors.NewInfraError(err, id)
	}
	return menu, nil
}
func (p menuPersistence) GetList(memoID int) ([]entity.Menu, error) {
	var menus []entity.Menu
	Db := database.Db
	err := Db.Where("memo_id = ?", memoID).Find(&menus).Error
	if err != nil {
		return menus, errors.NewInfraError(err, memoID)
	}
	return menus, nil
}
