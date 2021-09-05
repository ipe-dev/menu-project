package usecase

import (
	"log"

	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/domain/service"
)

type MenuUseCase interface {
	GetList(GetMenuListRequest) ([]entity.Menu, error)
	BulkCreate(BulkCreateMenuRequest) error
	BulkUpdate(BulkUpdateMenuRequest) error
	Get(GetMenuRequest) (entity.Menu, error)
}
type menuUseCase struct {
	menuRepository repository.MenuRepository
	memoRepository repository.MemoRepository
	memoService    service.MemoService
}

func NewMenuUseCase(menuRepo repository.MenuRepository, memoRepo repository.MemoRepository, s service.MemoService) MenuUseCase {
	return menuUseCase{
		menuRepository: menuRepo,
		memoRepository: memoRepo,
		memoService:    s,
	}
}

type CreateMenuRequest struct {
	Name string `json:"name"`
	Date int64  `json:"date" binding:"required"`
	Kind int    `json:"kind" binding:"required"`
	URL  string `json:"url"`
}
type UpdateMenuRequest struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name"`
	Date int64  `json:"date" binding:"required"`
	Kind int    `json:"kind" binding:"required"`
	URL  string `json:"url"`
}
type GetMenuListRequest struct {
	MemoID int `json:"memo_id" binding:"required"`
}
type GetMenuRequest struct {
	ID int `json:"id" binding:"required"`
}
type BulkCreateMenuRequest struct {
	MemoID         int `json:"memo_id"`
	CreateRequests []CreateMenuRequest
}
type BulkUpdateMenuRequest struct {
	MemoID         int `json:"memo_id"`
	UpdateRequests []UpdateMenuRequest
}

func (u menuUseCase) GetList(r GetMenuListRequest) ([]entity.Menu, error) {
	menus, err := u.menuRepository.GetList(r.MemoID)
	if err != nil {
		return menus, err
	}
	return menus, nil
}

func (u menuUseCase) BulkCreate(r BulkCreateMenuRequest) error {

	var menus []entity.Menu
	var err error

	// memoの存在チェック
	if err = u.memoService.CheckMemoExists(r.MemoID); err != nil {
		return err
	}

	memo, err := u.memoRepository.GetByID(r.MemoID)
	if err != nil {
		return err
	}

	for _, mr := range r.CreateRequests {
		menu := entity.NewMenu(
			entity.MenuNameOption(mr.Name),
			entity.MenuDateOption(mr.Date),
			entity.MenuKindOption(mr.Kind),
			entity.MenuUrlOption(mr.URL),
			entity.MenuMemoIDOption(r.MemoID),
		)
		if err = menu.CheckMenuDate(memo); err != nil {
			return err
		}

		menus = append(menus, *menu)
	}
	err = u.menuRepository.BulkCreate(menus)

	return err
}
func (u menuUseCase) BulkUpdate(r BulkUpdateMenuRequest) error {
	var menus []entity.Menu
	var err error
	if err = u.memoService.CheckMemoExists(r.MemoID); err != nil {
		return err
	}
	memo, err := u.memoRepository.GetByID(r.MemoID)
	if err != nil {
		return err
	}

	for _, mr := range r.UpdateRequests {
		menu := entity.NewMenu(
			entity.MenuIDOption(mr.ID),
			entity.MenuNameOption(mr.Name),
			entity.MenuDateOption(mr.Date),
			entity.MenuKindOption(mr.Kind),
			entity.MenuUrlOption(mr.URL),
			entity.MenuMemoIDOption(r.MemoID),
		)
		if err = menu.CheckMenuDate(memo); err != nil {
			return err
		}
		menus = append(menus, *menu)
	}
	err = u.menuRepository.BulkUpdate(menus)
	if err != nil {
		log.Println(err)
	}
	return err
}
func (u menuUseCase) Get(r GetMenuRequest) (entity.Menu, error) {
	var menu entity.Menu
	var err error
	menu, err = u.menuRepository.GetByID(r.ID)

	return menu, err

}
