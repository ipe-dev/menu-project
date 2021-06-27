package persistence

import (
	"time"

	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/infrastructure/database"
)

type menuPersistence struct{}

func NewMenuPersistence() repository.MenuRepository {
	return &menuPersistence{}
}
func (p menuPersistence) BulkCreate(menus []entity.Menu) ([]entity.Menu, error) {
	tx := database.Db.Begin()
	err := tx.Model(&menus).Create(menus).Error
	if err != nil {
		tx.Rollback()
		return menus, err
	}
	return menus, err
}
func (p menuPersistence) BulkUpdate(menus []entity.Menu) ([]entity.Menu, error) {
	tx := database.Db.Begin()
	var res []entity.Menu
	var err error
	for _, menu := range menus {
		err = tx.Model(&menu).Updates(menu).Error
		if err != nil {
			tx.Rollback()
			return res, err
		}
		res = append(res, menu)
	}
	tx.Commit()
	return res, err
}
func (p menuPersistence) GetByID(id int) (entity.Menu, error) {
	var menu entity.Menu
	Db := database.Db
	err := Db.Table("menus").Where("id = ?", id).First(&menu).Error
	return menu, err
}
func (p menuPersistence) GetByDate(date int64, userID int) (entity.Menu, error) {
	var menu entity.Menu
	Db := database.Db
	stringDate := time.Unix(date, 0).Format("2006/01/02 15:05:05")
	err := Db.Table("menus").Where("date = ?", stringDate).Where("user_id = ?", userID).First(&menu).Error
	return menu, err
}
func (p menuPersistence) GetList(weekID int, userID int) ([]entity.Menu, error) {
	var menus []entity.Menu
	Db := database.Db
	err := Db.Where("week_id = ?", weekID).Where("user_id = ?", userID).Find(&menus).Error
	return menus, err
}
