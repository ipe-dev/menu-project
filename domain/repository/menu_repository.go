package repository

import (
	"github.com/ipe-dev/menu_project/domain/entity"
)

type MenuRepository interface {
	Create(m *entity.Menu) error
	Update(m *entity.Menu) error
	Get(id int) (*entity.Menu, error)
	GetList(weekID int) ([]entity.Menu, error)
	Delete(id int) error
}
