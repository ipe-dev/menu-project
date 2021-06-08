package persistance

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/infrastructure/database"
)

type subMenuPersistance struct{}

func NewSubMenuPersistance() repository.SubMenuRepository {
	return &subMenuPersistance{}
}
func (p subMenuPersistance) GetByMenuID(MenuID int) ([]entity.SubMenu, error) {
	var submenus []entity.SubMenu
	Db := database.Db
	err := Db.Table("sub_menus").Where("menu_id = ?", MenuID).Find(&submenus).Error
	return submenus, err
}
func (p subMenuPersistance) GetList(MenuIDList []int) ([]entity.SubMenu, error) {
	var submenus []entity.SubMenu
	Db := database.Db
	err := Db.Where("menu_id IN ?", MenuIDList).Find(&submenus).Error
	return submenus, err
}

func (p subMenuPersistance) BulkCreate(submenus []entity.SubMenu) ([]entity.SubMenu, error) {
	Db := database.Db
	tx := Db.Begin()
	err := tx.Create(&submenus).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return submenus, err
}
func (p subMenuPersistance) BulkUpdate(submenus []entity.SubMenu) ([]entity.SubMenu, error) {
	Db := database.Db
	tx := Db.Begin()
	var err error
	for _, v := range submenus {
		err = tx.Updates(v).Error
		if err != nil {
			tx.Rollback()
			return submenus, err
		}
	}
	tx.Commit()
	return submenus, err
}
