package usecase

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
)

type SubMenuUseCase interface {
	GetSubMenu(GetSubMenuRequest) ([]entity.SubMenu, error)
	GetSubMenuList(GetSubMenuListRequest) ([]entity.SubMenu, error)
	BulkCreateSubMenu(BulkCreateSubMenuRequest) ([]entity.SubMenu, error)
	BulkUpdateSubMenu(BulkUpdateSubMenuRequest) ([]entity.SubMenu, error)
}
type subMenuUseCase struct {
	subMenuRepository repository.SubMenuRepository
}

func NewSubMenuUseCase(r repository.SubMenuRepository) SubMenuUseCase {
	return &subMenuUseCase{r}
}

type GetSubMenuRequest struct {
	ID     int `json:"id"`
	MenuID int `json:"menu_id"`
}
type GetSubMenuListRequest struct {
	MenuIDList []int `json:"menu_id_list"`
}
type CreateSubMenuRequest struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	MenuID int    `json:"menu_id"`
}
type BulkCreateSubMenuRequest struct {
	CreateRequests []CreateSubMenuRequest
}
type UpdateSubMenuRequest struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	MenuID int    `json:"menu_id"`
}
type BulkUpdateSubMenuRequest struct {
	Updaterequests []UpdateSubMenuRequest
}

func (su subMenuUseCase) GetSubMenu(r GetSubMenuRequest) ([]entity.SubMenu, error) {
	var submenus []entity.SubMenu
	var err error
	if r.ID != 0 {
		submenus, err = su.subMenuRepository.GetByMenuID(r.MenuID)
	}
	return submenus, err
}

func (su subMenuUseCase) GetSubMenuList(r GetSubMenuListRequest) ([]entity.SubMenu, error) {
	submenus, err := su.subMenuRepository.GetList(r.MenuIDList)
	return submenus, err
}
func (su subMenuUseCase) BulkCreateSubMenu(bc BulkCreateSubMenuRequest) ([]entity.SubMenu, error) {
	var submenus []entity.SubMenu
	for _, v := range bc.CreateRequests {
		submenu := entity.SubMenu{
			Name:   v.Name,
			MenuID: v.MenuID,
		}
		submenus = append(submenus, submenu)
	}
	submenus, err := su.subMenuRepository.BulkCreate(submenus)
	return submenus, err
}
func (su subMenuUseCase) BulkUpdateSubMenu(bu BulkUpdateSubMenuRequest) ([]entity.SubMenu, error) {
	var submenus []entity.SubMenu
	for _, v := range bu.Updaterequests {
		submenu := entity.SubMenu{
			ID:     v.ID,
			Name:   v.Name,
			MenuID: v.MenuID,
		}
		submenus = append(submenus, submenu)
	}
	submenus, err := su.subMenuRepository.BulkUpdate(submenus)
	return submenus, err
}
