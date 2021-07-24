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
func (p foodStuffPersistence) BulkCreate(FoodStuffs []entity.FoodStuff) ([]entity.FoodStuff, error) {
	tx := database.Db.Begin()
	err := tx.Model(&FoodStuffs).Create(FoodStuffs).Error
	if err != nil {
		tx.Rollback()
		return FoodStuffs, errors.NewInfraError(err, FoodStuffs)
	}
	tx.Commit()
	return FoodStuffs, nil
}
func (p foodStuffPersistence) BulkUpdate(FoodStuffs []entity.FoodStuff) ([]entity.FoodStuff, error) {
	tx := database.Db.Begin()
	var res []entity.FoodStuff
	var err error
	for _, s := range FoodStuffs {
		err = tx.Model(&s).Updates(s).Error
		if err != nil {
			tx.Rollback()
			return res, errors.NewInfraError(err, s)
		}
		res = append(res, s)
	}
	tx.Commit()
	return res, nil
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
