package repository

import "github.com/ipe-dev/menu_project/domain/entity"

type SubMenuRepository interface {
	GetByMemoID(MemoID int) ([]entity.SubMenu, error)
	GetList(MenuIDList []int) ([]entity.SubMenu, error)
	BulkCreate(SubMenus []entity.SubMenu) ([]entity.SubMenu, error)
	BulkUpdate(SubMenus []entity.SubMenu) ([]entity.SubMenu, error)
}
