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
func (p menuPersistence) BulkCreate(menus []entity.Menu) ([]entity.Menu, error) {
	tx := database.Db.Begin()
	err := tx.Model(&menus).Create(menus).Error
	if err != nil {
		tx.Rollback()
		return menus, errors.NewInfraError(err, menus)
	}
	tx.Commit()
	return menus, nil
}
func (p menuPersistence) BulkUpdate(menus []entity.Menu) ([]entity.Menu, error) {
	tx := database.Db.Begin()
	var res []entity.Menu
	var err error
	for _, menu := range menus {
		err = tx.Model(&menu).Updates(menu).Error
		if err != nil {
			tx.Rollback()
			return res, errors.NewInfraError(err, menu)
		}
		res = append(res, menu)
	}
	tx.Commit()
	return res, nil
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
