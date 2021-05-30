package repository

import "github.com/ipe-dev/menu_project/domain/entity"

type SubMenuRepository interface {
	GetByMenuID(MenuID int) ([]entity.SubMenu, error)
	GetList(MenuIDList []int) ([]entity.SubMenu, error)
	BulkCreate([]entity.SubMenu) ([]entity.SubMenu, error)
	BulkUpdate([]entity.SubMenu) ([]entity.SubMenu, error)
}
