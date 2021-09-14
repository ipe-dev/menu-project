package repository

import "github.com/ipe-dev/menu_project/domain/entity"

type FoodStuffRepository interface {
	GetByMenuID(MenuID int) (entity.FoodStuff, error)
	GetList(MenuIDList []int) ([]entity.FoodStuff, error)
	Save([]entity.FoodStuff) error
	ChangeBuyStatus(ID int, Status int) error
}
