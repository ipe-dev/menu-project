package repository

import "github.com/ipe-dev/menu_project/domain/entity"

type SubFoodStuffRepository interface {
	GetBySubMenuID(SubMenuID int) (entity.SubFoodStuff, error)
	GetList(SubMenuIDList []int) ([]entity.SubFoodStuff, error)
	BulkCreate([]entity.SubFoodStuff) ([]entity.SubFoodStuff, error)
	BulkUpdate([]entity.SubFoodStuff) ([]entity.SubFoodStuff, error)
	ChangeBuyStatus(ID int, Status int) error
}
