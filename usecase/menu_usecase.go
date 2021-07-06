package usecase

import (
	"log"

	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/infrastructure/factory"
)

type MenuUseCase interface {
	GetList(GetMenuListRequest) ([]entity.Menu, error)
	BulkCreate(BulkCreateMenuRequest) ([]entity.Menu, error)
	BulkUpdate(BulkUpdateMenuRequest) ([]entity.Menu, error)
	Get(GetMenuRequest) (entity.Menu, error)
}
type menuUseCase struct {
	menuRepository repository.MenuRepository
	weekIDFactory  factory.WeekIDFactory
}

func NewMenuUseCase(r repository.MenuRepository, f factory.WeekIDFactory) MenuUseCase {
	return menuUseCase{
		menuRepository: r,
		weekIDFactory:  f,
	}
}

type CreateMenuRequest struct {
	Name   string `json:"name"`
	Date   int64  `json:"date" validate:"required"`
	Kind   int    `json:"kind" validate:"required"`
	URL    string `json:"url"`
	UserID int    `json:"user_id" validate:"required"`
	WeekID int    `json:"week_id" validate:"required"`
}
type UpdateMenuRequest struct {
	ID     int    `json:"id" validate:"required"`
	Name   string `json:"name"`
	Date   int64  `json:"date" validate:"required"`
	Kind   int    `json:"kind" validate:"required"`
	URL    string `json:"url"`
	UserID int    `json:"user_id" validate:"required"`
}
type GetMenuListRequest struct {
	WeekID int `json:"week_id" validate:"required"`
	UserID int `json:"user_id" validate:"required"`
}
type GetMenuRequest struct {
	ID     int   `json:"id" validate:"required"`
	Date   int64 `json:"date" validate:"required"`
	UserID int   `json:"user_id" validate:"required"`
}
type BulkCreateMenuRequest struct {
	CreateRequests []CreateMenuRequest
}
type BulkUpdateMenuRequest struct {
	UpdateRequests []UpdateMenuRequest
}

func (u menuUseCase) GetList(r GetMenuListRequest) ([]entity.Menu, error) {
	menus, err := u.menuRepository.GetList(r.WeekID, r.UserID)
	return menus, err
}

func (u menuUseCase) BulkCreate(r BulkCreateMenuRequest) ([]entity.Menu, error) {

	var menus []entity.Menu
	WeekID, err := u.weekIDFactory.NewWeekID(r.CreateRequests[0].UserID)
	if err != nil {
		return menus, err
	}
	for _, mr := range r.CreateRequests {
		menu := entity.NewMenu(
			entity.MenuNameOption(mr.Name),
			entity.MenuDateOption(mr.Date),
			entity.MenuKindOption(mr.Kind),
			entity.MenuUrlOption(mr.URL),
			entity.MenuUserIDOption(mr.UserID),
			entity.MenuWeekIDOption(WeekID),
		)
		menus = append(menus, *menu)
	}
	menus, err = u.menuRepository.BulkCreate(menus)
	if err != nil {
		return menus, err
	}
	err = u.weekIDFactory.IncrementWeekID(r.CreateRequests[0].UserID)
	if err != nil {
		return menus, err
	}

	return menus, err
}
func (u menuUseCase) BulkUpdate(r BulkUpdateMenuRequest) ([]entity.Menu, error) {
	var menus []entity.Menu
	for _, mr := range r.UpdateRequests {
		menu := entity.NewMenu(
			entity.MenuIDOption(mr.ID),
			entity.MenuNameOption(mr.Name),
			entity.MenuDateOption(mr.Date),
			entity.MenuKindOption(mr.Kind),
			entity.MenuUrlOption(mr.URL),
			entity.MenuUserIDOption(mr.UserID),
		)
		menus = append(menus, *menu)
	}
	menus, err := u.menuRepository.BulkUpdate(menus)
	if err != nil {
		log.Println(err)
	}
	return menus, err
}
func (u menuUseCase) Get(r GetMenuRequest) (entity.Menu, error) {
	var menu entity.Menu
	var err error
	if r.ID != 0 {
		menu, err = u.menuRepository.GetByID(r.ID)
	} else if r.Date != 0 {
		menu, err = u.menuRepository.GetByDate(r.Date, r.UserID)
	}
	return menu, err

}
