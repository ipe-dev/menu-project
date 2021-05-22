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

func (m menuPersistance) Create(menu *entity.Menu) error {
	Db := database.Db
	menu.CreatedAt = time.Now().Format("2006/01/02 15:05:05")
	err := Db.Create(menu).Error
	return err
}
func (m menuPersistance) Update(menu *entity.Menu) error {
	Db := database.Db
	menu.UpdatedAt = time.Now().Format("2006/01/02 15:05:05")
	err := Db.Model(&menu).Updates(menu).Error
	return err
}
func (m menuPersistance) Get(id int) (*entity.Menu, error) {
	var menu entity.Menu
	Db := database.Db
	err := Db.Table("menus").Where("id = ?", id).First(&menu).Error
	return &menu, err
}
func (m menuPersistance) GetList(weekID int) ([]entity.Menu, error) {
	var menus []entity.Menu
	Db := database.Db
	err := Db.Where("week_id = ?", weekID).Find(menus).Error
	return menus, err
}
func (m menuPersistance) Delete(id int) error {
	Db := database.Db
	err := Db.Table("menus").Where("id = ?", id).Delete().Error
	return err
}
