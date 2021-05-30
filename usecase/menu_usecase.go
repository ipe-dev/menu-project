package usecase

import (
	"log"
	"time"

	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
)

type MenuUseCase interface {
	GetMenuList(GetMenuListRequest) ([]entity.Menu, error)
	BulkCreateMenu(BulkCreateMenuRequest) ([]entity.Menu, error)
	BulkUpdateMenu(BulkUpdateMenuRequest) ([]entity.Menu, error)
	GetMenu(GetMenuRequest) (*entity.Menu, error)
}
type menuUseCase struct {
	menuRepository repository.MenuRepository
}

func NewMenuUseCase(r repository.MenuRepository) MenuUseCase {
	return &menuUseCase{r}
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
	ID   int   `json:"id"`
	Date int64 `json:"date"`
}
type BulkCreateMenuRequest struct {
	CreateRequests []CreateMenuRequest
}
type BulkUpdateMenuRequest struct {
	UpdateRequests []UpdateMenuRequest
}

func (u menuUseCase) BulkCreateMenu(bc BulkCreateMenuRequest) ([]entity.Menu, error) {

	var menus []entity.Menu
	for _, mr := range bc.CreateRequests {
		menu := entity.Menu{
			Name:   mr.Name,
			Date:   time.Unix(mr.Date, 0).Format("2006/01/02 15:05:05"),
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
func (u menuUseCase) BulkUpdateMenu(r BulkUpdateMenuRequest) ([]entity.Menu, error) {
	var menus []entity.Menu
	for _, mr := range r.UpdateRequests {
		menu := entity.Menu{
			ID:     mr.ID,
			Name:   mr.Name,
			Date:   time.Unix(mr.Date, 0).Format("2006/01/02 15:05:05"),
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
func (u menuUseCase) GetMenu(gr GetMenuRequest) (*entity.Menu, error) {
	var menu *entity.Menu
	var err error
	if gr.ID != 0 {
		menu, err = u.menuRepository.GetByID(gr.ID)
		if err != nil {
			log.Println(err)
		}
	} else if gr.Date != 0 {
		menu, err = u.menuRepository.GetByDate(gr.Date, gr.ID)
	}
	return menu, err

}
func (u menuUseCase) GetMenuList(gr GetMenuListRequest) ([]entity.Menu, error) {
	menus, err := u.menuRepository.GetList(gr.WeekID, gr.UserID)
	if err != nil {
		log.Println(err)
	}
	return menus, err
}
