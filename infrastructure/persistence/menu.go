package persistance

import (
	"time"

	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/infrastructure/database"
)

type menuPersistance struct{}

func NewMenuPersistance() repository.MenuRepository {
	return &menuPersistance{}
}
func (m menuPersistance) BulkCreate(menus []entity.Menu) ([]entity.Menu, error) {
	Db := database.Db
	err := Db.Create(&menus).Error
	return menus, err
}
func (m menuPersistance) BulkUpdate(menus []entity.Menu) ([]entity.Menu, error) {
	Db := database.Db
	err := Db.Updates(&menus).Error
	return menus, err
}
func (m menuPersistance) GetByID(id int) (*entity.Menu, error) {
	var menu entity.Menu
	Db := database.Db
	err := Db.Table("menus").Where("id = ?", id).First(&menu).Error
	return &menu, err
}
func (m menuPersistance) GetByDate(date int64, userID int) (*entity.Menu, error) {
	var menu entity.Menu
	Db := database.Db
	stringDate := time.Unix(date, 0).Format("2006/01/02 15:05:05")
	err := Db.Table("menus").Where("date = ?", stringDate).Where("user_id = ?", userID).First(&menu).Error
	return &menu, err
}
func (m menuPersistance) GetList(weekID int, userID int) ([]entity.Menu, error) {
	var menus []entity.Menu
	Db := database.Db
	err := Db.Where("week_id = ?", weekID).Where("user_id = ?", userID).Find(menus).Error
	return menus, err
}
