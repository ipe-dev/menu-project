package persistence

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/errors"
	"github.com/ipe-dev/menu_project/infrastructure/database"
)

type subFoodStuffPersistence struct{}

func NewSubFoodStuffPersistence() repository.SubFoodStuffRepository {
	return &subFoodStuffPersistence{}
}
func (p subFoodStuffPersistence) GetBySubMenuID(MenuID int) (entity.SubFoodStuff, error) {
	var SubFoodStuff entity.SubFoodStuff
	Db := database.Db
	err := Db.Table("sub_food_stuffs").Where("sub_menu_id = ?", MenuID).First(&SubFoodStuff).Error
	if err != nil {
		return SubFoodStuff, errors.NewInfraError(err, MenuID)
	}
	return SubFoodStuff, nil
}
func (p subFoodStuffPersistence) GetList(MenuIDList []int) ([]entity.SubFoodStuff, error) {
	var SubFoodStuffs []entity.SubFoodStuff
	Db := database.Db
	err := Db.Where("sub_menu_id IN ?", MenuIDList).Find(&SubFoodStuffs).Error
	if err != nil {
		return SubFoodStuffs, errors.NewInfraError(err, MenuIDList)
	}
	return SubFoodStuffs, nil
}
func (p subFoodStuffPersistence) BulkCreate(SubFoodStuffs []entity.SubFoodStuff) ([]entity.SubFoodStuff, error) {
	tx := database.Db.Begin()
	err := tx.Model(&SubFoodStuffs).Create(SubFoodStuffs).Error
	if err != nil {
		tx.Rollback()
		return SubFoodStuffs, errors.NewInfraError(err, SubFoodStuffs)
	}
	tx.Commit()
	return SubFoodStuffs, nil
}
func (p subFoodStuffPersistence) BulkUpdate(SubFoodStuffs []entity.SubFoodStuff) ([]entity.SubFoodStuff, error) {
	tx := database.Db.Begin()
	var res []entity.SubFoodStuff
	var err error
	for _, s := range SubFoodStuffs {
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
func (p subFoodStuffPersistence) ChangeBuyStatus(ID int, Status int) error {
	tx := database.Db.Begin()
	err := tx.Table("sub_food_stuffs").Where("id = ?", ID).Update("buy_status", Status).Error
	if err != nil {
		tx.Rollback()
		return errors.NewInfraError(err, ID, Status)
	}
	tx.Commit()
	return nil
}
