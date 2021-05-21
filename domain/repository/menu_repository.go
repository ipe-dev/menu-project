package repository

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/usecase"
)

type MenuRepository interface {
	Create(m *entity.Menu) error
	Update(m *entity.Menu) error
	Get(r usecase.GetMenuRequest) (*entity.Menu, error)
	GetList(r usecase.GetMenuListRequest) ([]*entity.Menu, error)
	Delete(id int) error
}
