package usecase

import (
	"log"

	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
)

type MenuUseCase interface {
	GetList(GetMenuListRequest) ([]entity.Menu, error)
	BulkCreate(BulkCreateMenuRequest) ([]entity.Menu, error)
	BulkUpdate(BulkUpdateMenuRequest) ([]entity.Menu, error)
	Get(GetMenuRequest) (entity.Menu, error)
}
type menuUseCase struct {
	menuRepository repository.MenuRepository
}

func NewMenuUseCase(r repository.MenuRepository) MenuUseCase {
	return menuUseCase{
		menuRepository: r,
	}
}

type CreateMenuRequest struct {
	Name   string `json:"name"`
	Date   int64  `json:"date" validate:"required"`
	Kind   int    `json:"kind" validate:"required"`
	URL    string `json:"url"`
	UserID int    `json:"user_id" validate:"required"`
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
	UserID int `json:"user_id" validate:"required"`
	MemoID int `json:"memo_id" validate:"required"`
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
	menus, err := u.menuRepository.GetList(r.MemoID, r.UserID)
	if err != nil {
		return menus, err
	}
	return menus, nil
}

func (u menuUseCase) BulkCreate(r BulkCreateMenuRequest) ([]entity.Menu, error) {

	var menus []entity.Menu

	for _, mr := range r.CreateRequests {
		menu := entity.NewMenu(
			entity.MenuNameOption(mr.Name),
			entity.MenuDateOption(mr.Date),
			entity.MenuKindOption(mr.Kind),
			entity.MenuUrlOption(mr.URL),
			entity.MenuUserIDOption(mr.UserID),
		)
		menus = append(menus, *menu)
	}
	menus, err := u.menuRepository.BulkCreate(menus)

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
