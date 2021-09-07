package repository

import (
	"github.com/ipe-dev/menu_project/domain/entity"
)

type MemoRepository interface {
	Create(memo entity.Memo) error
	Update(memo entity.Memo) error
	GetByID(id int) (entity.Memo, error)
	GetList(userID int) ([]entity.Memo, error)
}
