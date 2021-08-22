package repository

import (
	"github.com/ipe-dev/menu_project/domain/entity"
)

type MenuRepository interface {
	BulkCreate(menus []entity.Menu) ([]entity.Menu, error)
	BulkUpdate(menus []entity.Menu) ([]entity.Menu, error)
	GetByID(id int) (entity.Menu, error)
	GetList(memoID int) ([]entity.Menu, error)
}
