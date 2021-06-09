package persistence

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/infrastructure/database"
)

type foodStuffPersistence struct{}

func NewFoodStuffPersistence() repository.FoodStuffRepository {
	return &foodStuffPersistence{}
}
func (p foodStuffPersistence) GetByMenuID(MenuID int) (entity.FoodStuff, error) {
	var FoodStuff entity.FoodStuff
	Db := database.Db
	err := Db.Table("food_stuffs").Where("menu_id = ?", MenuID).First(&FoodStuff).Error
	return FoodStuff, err
}
func (p foodStuffPersistence) GetList(MenuIDList []int) ([]entity.FoodStuff, error) {
	var FoodStuffs []entity.FoodStuff
	Db := database.Db
	err := Db.Where("menu_id IN ?", MenuIDList).Find(&FoodStuffs).Error
	return FoodStuffs, err
}
func (p foodStuffPersistence) BulkCreate(FoodStuffs []entity.FoodStuff) ([]entity.FoodStuff, error) {
	Db := database.Db
	err := Db.Model(&FoodStuffs).Create(FoodStuffs).Error
	return FoodStuffs, err
}
func (p foodStuffPersistence) BulkUpdate(FoodStuffs []entity.FoodStuff) ([]entity.FoodStuff, error) {
	Db := database.Db
	var res []entity.FoodStuff
	var err error
	for _, s := range FoodStuffs {
		err = Db.Model(&s).Updates(s).Error
		res = append(res, s)
	}
	return res, err
}
