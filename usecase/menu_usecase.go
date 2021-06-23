package usecase

import (
	"log"
	"time"

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
	return menuUseCase{menuRepository: r}
}

type CreateMenuRequest struct {
	Name   string `json:"name"`
	Date   int64  `json:"date"`
	Kind   int    `json:"kind"`
	URL    string `json:"url"`
	UserID int    `json:"user_id"`
	WeekID int    `json:"week_id"`
}
type UpdateMenuRequest struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Date   int64  `json:"date"`
	Kind   int    `json:"kind"`
	URL    string `json:"url"`
	UserID int    `json:"user_id"`
	WeekID int    `json:"week_id"`
}
type GetMenuListRequest struct {
	WeekID int `json:"week_id"`
	UserID int `json:"user_id"`
}
type GetMenuRequest struct {
	ID     int   `json:"id"`
	Date   int64 `json:"date"`
	UserID int   `json:"user_id"`
}
type BulkCreateMenuRequest struct {
	CreateRequests []CreateMenuRequest
}
type BulkUpdateMenuRequest struct {
	UpdateRequests []UpdateMenuRequest
}

func (u menuUseCase) GetList(r GetMenuListRequest) ([]entity.Menu, error) {
	menus, err := u.menuRepository.GetList(r.WeekID, r.UserID)
	if err != nil {
		log.Println(err)
	}
	return menus, err
}

func (u menuUseCase) BulkCreate(r BulkCreateMenuRequest) ([]entity.Menu, error) {

	var menus []entity.Menu
	for _, mr := range r.CreateRequests {
		menu := entity.Menu{
			Name:   mr.Name,
			Date:   time.Unix(mr.Date, 0).Format("2006/01/02"),
			Kind:   mr.Kind,
			URL:    mr.URL,
			UserID: mr.UserID,
			WeekID: mr.WeekID,
		}
		menus = append(menus, menu)
	}
	menus, err := u.menuRepository.BulkCreate(menus)
	if err != nil {
		log.Println(err)
	}
	return menus, err
}
func (u menuUseCase) BulkUpdate(r BulkUpdateMenuRequest) ([]entity.Menu, error) {
	var menus []entity.Menu
	for _, mr := range r.UpdateRequests {
		menu := entity.Menu{
			ID:     mr.ID,
			Name:   mr.Name,
			Kind:   mr.Kind,
			URL:    mr.URL,
			UserID: mr.UserID,
			WeekID: mr.WeekID,
		}
		menus = append(menus, menu)
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
		if err != nil {
			log.Println(err)
		}
	} else if r.Date != 0 {
		menu, err = u.menuRepository.GetByDate(r.Date, r.UserID)
	}
	return menu, err

}
