package usecase

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
)

type SubMenuUseCase interface {
	Get(GetSubMenuRequest) ([]entity.SubMenu, error)
	GetList(GetSubMenuListRequest) ([]entity.SubMenu, error)
	BulkCreate(BulkCreateSubMenuRequest) ([]entity.SubMenu, error)
	BulkUpdate(BulkUpdateSubMenuRequest) ([]entity.SubMenu, error)
}
type subMenuUseCase struct {
	subMenuRepository repository.SubMenuRepository
}

func NewSubMenuUseCase(r repository.SubMenuRepository) SubMenuUseCase {
	return &subMenuUseCase{r}
}

type GetSubMenuRequest struct {
	ID     int `json:"id"　validate:"required"`
	MenuID int `json:"menu_id"　validate:"required"`
}
type GetSubMenuListRequest struct {
	MenuIDList []int `json:"menu_id_list"`
}
type CreateSubMenuRequest struct {
	ID     int    `json:"id"　validate:"required"`
	Name   string `json:"name"　validate:"required"`
	MenuID int    `json:"menu_id"　validate:"required"`
}
type BulkCreateSubMenuRequest struct {
	CreateRequests []CreateSubMenuRequest
}
type UpdateSubMenuRequest struct {
	ID     int    `json:"id"　validate:"required"`
	Name   string `json:"name"`
	MenuID int    `json:"menu_id"　validate:"required"`
}
type BulkUpdateSubMenuRequest struct {
	UpdateRequests []UpdateSubMenuRequest
}

func (u subMenuUseCase) Get(r GetSubMenuRequest) ([]entity.SubMenu, error) {
	var submenus []entity.SubMenu
	var err error
	if r.ID != 0 {
		submenus, err = u.subMenuRepository.GetByMenuID(r.MenuID)
	}
	return submenus, err
}

func (u subMenuUseCase) GetList(r GetSubMenuListRequest) ([]entity.SubMenu, error) {
	submenus, err := u.subMenuRepository.GetList(r.MenuIDList)
	return submenus, err
}
func (u subMenuUseCase) BulkCreate(bc BulkCreateSubMenuRequest) ([]entity.SubMenu, error) {
	var submenus []entity.SubMenu
	for _, v := range bc.CreateRequests {
		submenu := entity.NewSubMenu(
			entity.SubMenuNameOption(v.Name),
			entity.SubMenuMenuIDOption(v.MenuID),
		)
		submenus = append(submenus, *submenu)
	}
	submenus, err := u.subMenuRepository.BulkCreate(submenus)
	return submenus, err
}
func (u subMenuUseCase) BulkUpdate(r BulkUpdateSubMenuRequest) ([]entity.SubMenu, error) {
	var submenus []entity.SubMenu
	for _, v := range r.UpdateRequests {
		submenu := entity.NewSubMenu(
			entity.SubMenuIDOption(v.ID),
			entity.SubMenuNameOption(v.Name),
			entity.SubMenuMenuIDOption(v.MenuID),
		)

		submenus = append(submenus, *submenu)
	}
	submenus, err := u.subMenuRepository.BulkUpdate(submenus)
	return submenus, err
}
