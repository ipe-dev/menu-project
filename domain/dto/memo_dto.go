package dto

import "github.com/ipe-dev/menu_project/domain/entity"

type MemoDto struct {
	entity.Memo
	Menus      []entity.Menu
	SubMenus   []entity.SubMenu
	FoodStuffs []entity.FoodStuff
}
