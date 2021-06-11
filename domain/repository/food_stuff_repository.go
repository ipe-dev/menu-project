package repository

import "github.com/ipe-dev/menu_project/domain/entity"

type FoodStuffRepository interface {
	GetByMenuID(MenuID int) (entity.FoodStuff, error)
	GetList(MenuIDList []int) ([]entity.FoodStuff, error)
	BulkCreate([]entity.FoodStuff) ([]entity.FoodStuff, error)
	BulkUpdate([]entity.FoodStuff) ([]entity.FoodStuff, error)
	ChangeBuyStatus(ID int, Status int) error
}
