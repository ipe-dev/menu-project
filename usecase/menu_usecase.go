package usecase

import (
	"time"

	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
)

type MenuUseCase interface {
	BulkCreateMenu(BulkCreateMenuRequest) ([]entity.Menu, error)
	BulkUpdateMenu(BulkUpdateMenuRequest) ([]entity.Menu, error)
	GetMenuByID(GetMenuRequest) (*entity.Menu, error)
	GetMenuByDate(GetMenuRequest) (*entity.Menu, error)
}
type menuUseCase struct {
	menuRepository repository.MenuRepository
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

func NewMenuUseCase(r repository.MenuRepository) MenuUseCase {
	return &menuUseCase{r}
}

func (mu menuUseCase) BulkCreateMenu(bc BulkCreateMenuRequest) ([]entity.Menu, error) {

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
	menus, err := mu.menuRepository.BulkCreate(menus)
	return menus, err
}
func (mu menuUseCase) BulkUpdateMenu(bu BulkUpdateMenuRequest) ([]entity.Menu, error) {
	var menus []entity.Menu
	for _, mr := range bu.UpdateRequests {
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
	menus, err := mu.menuRepository.BulkUpdate(menus)
	return menus, err
}

func (mu menuUseCase) GetMenuByID(gr GetMenuRequest) (*entity.Menu, error) {
	menu, err := mu.menuRepository.GetByID(gr.ID)
	return menu, err
}
func (mu menuUseCase) GetMenuByDate(gr GetMenuRequest) (*entity.Menu, error) {
	menu, err := mu.menuRepository.GetByDate(gr.Date, gr.ID)
	return menu, err
}
func (mu menuUseCase) GetMenuList(gr GetMenuListRequest) ([]entity.Menu, error) {
	menus, err := mu.menuRepository.GetList(gr.WeekID, gr.UserID)
	return menus, err
}
