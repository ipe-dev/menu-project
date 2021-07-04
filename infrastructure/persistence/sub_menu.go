package persistence

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/errors"
	"github.com/ipe-dev/menu_project/infrastructure/database"
)

type subMenuPersistence struct{}

func NewSubMenuPersistence() repository.SubMenuRepository {
	return &subMenuPersistence{}
}
func (p subMenuPersistence) GetByMenuID(MenuID int) ([]entity.SubMenu, error) {
	var submenus []entity.SubMenu
	Db := database.Db
	err := Db.Table("sub_menus").Where("menu_id = ?", MenuID).Find(&submenus).Error
	return submenus, errors.NewInfraError(err, MenuID)
}
func (p subMenuPersistence) GetList(MenuIDList []int) ([]entity.SubMenu, error) {
	var submenus []entity.SubMenu
	Db := database.Db
	err := Db.Where("menu_id IN ?", MenuIDList).Find(&submenus).Error
	return submenus, errors.NewInfraError(err, MenuIDList)
}

func (p subMenuPersistence) BulkCreate(submenus []entity.SubMenu) ([]entity.SubMenu, error) {
	tx := database.Db.Begin()
	err := tx.Create(&submenus).Error
	if err != nil {
		tx.Rollback()
		return submenus, errors.NewInfraError(err, submenus)
	}
	tx.Commit()
	return submenus, nil
}
func (p subMenuPersistence) BulkUpdate(submenus []entity.SubMenu) ([]entity.SubMenu, error) {
	tx := database.Db.Begin()
	var err error
	for _, v := range submenus {
		err = tx.Updates(v).Error
		if err != nil {
			tx.Rollback()
			return submenus, errors.NewInfraError(err, v)
		}
	}
	tx.Commit()
	return submenus, nil
}
