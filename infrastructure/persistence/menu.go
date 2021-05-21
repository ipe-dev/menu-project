package persistance

import (
	"time"

	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/infrastructure/database"
	"github.com/sqs/goreturns/returns"
)

type menuPersistance struct{}

var Db database.Db

func NewMenuPersistance() repository.MenuRepository {
	return &menuPersistance{}
}

func (m menuPersistance) Create(menu *entity.Menu) error {
	menu.CreatedAt = time.Now().Format("2006/01/02 15:05:05")
	err := Db.Create(menu).Error
	return err
}
func (m menuPersistance) Update(menu *entity.Menu) error {
	menu.UpdatedAt = time.Now().Format("2006/01/02 15:05:05")
	return
}
func (m menuPersistance) Get() error {
	return
}
func (m menuPersistance) GetList(*entity.Menu) error {
	return
}
func (m menuPersistance) Delete(id int) error {
	return
}
