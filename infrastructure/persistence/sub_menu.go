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
func (p subMenuPersistence) GetByMemoID(MemoID int) ([]entity.SubMenu, error) {
	var submenus []entity.SubMenu
	Db := database.Db
	err := Db.Table("sub_menus").Where("memo_id = ?", MemoID).Find(&submenus).Error
	if err != nil {
		return submenus, errors.NewInfraError(err, MemoID)
	}
	return submenus, nil
}
func (p subMenuPersistence) GetList(MenuIDList []int) ([]entity.SubMenu, error) {
	var submenus []entity.SubMenu
	Db := database.Db
	err := Db.Where("menu_id IN ?", MenuIDList).Find(&submenus).Error
	if err != nil {
		return submenus, errors.NewInfraError(err, MenuIDList)
	}
	return submenus, nil
}

func (p subMenuPersistence) Save(submenus []entity.SubMenu) error {
	tx := database.Db.Begin()
	var err error
	for _, submenu := range submenus {
		err = tx.Where("id = ?", submenu.ID).Assign(submenu).FirstOrCreate(&submenu).Error
		if err != nil {
			tx.Rollback()
			return errors.NewInfraError(err, submenu)
		}
	}
	tx.Commit()
	return nil
}
