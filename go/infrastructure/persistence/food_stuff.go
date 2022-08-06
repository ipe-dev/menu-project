package persistence

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/errors"
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
	if err != nil {
		return FoodStuff, errors.NewInfraError(err, MenuID)
	}
	return FoodStuff, nil
}
func (p foodStuffPersistence) GetList(MenuIDList []int) ([]entity.FoodStuff, error) {
	var FoodStuffs []entity.FoodStuff
	Db := database.Db
	err := Db.Where("menu_id IN ?", MenuIDList).Find(&FoodStuffs).Error
	if err != nil {
		return FoodStuffs, errors.NewInfraError(err, MenuIDList)
	}
	return FoodStuffs, nil
}

func (p foodStuffPersistence) Save(FoodStuffs []entity.FoodStuff) error {
	tx := database.Db.Begin()
	var err error
	for _, foodstuff := range FoodStuffs {
		err = tx.Where("id = ?", foodstuff.ID).Assign(foodstuff).FirstOrCreate(&foodstuff).Error
		if err != nil {
			tx.Rollback()
			return errors.NewInfraError(err, foodstuff)
		}
	}
	tx.Commit()
	return nil
}
func (p foodStuffPersistence) ChangeBuyStatus(ID int, Status int) error {
	tx := database.Db.Begin()
	err := tx.Table("food_stuffs").Where("id = ?", ID).Update("buy_status", Status).Error
	if err != nil {
		tx.Rollback()
		return errors.NewInfraError(err, ID, Status)
	}
	tx.Commit()
	return nil
}
