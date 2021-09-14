package usecase

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/usecase/requests"
)

type SubMenuUseCase interface {
	Get(requests.GetSubMenuRequest) ([]entity.SubMenu, error)
	GetList(requests.GetSubMenuListRequest) ([]entity.SubMenu, error)
	BulkCreate(requests.BulkCreateSubMenuRequest) error
}
type subMenuUseCase struct {
	subMenuRepository repository.SubMenuRepository
}

func NewSubMenuUseCase(r repository.SubMenuRepository) SubMenuUseCase {
	return &subMenuUseCase{r}
}

func (u subMenuUseCase) Get(r requests.GetSubMenuRequest) ([]entity.SubMenu, error) {
	var submenus []entity.SubMenu
	var err error
	if r.ID != 0 {
		submenus, err = u.subMenuRepository.GetByMemoID(r.MemoID)
	}
	return submenus, err
}

func (u subMenuUseCase) GetList(r requests.GetSubMenuListRequest) ([]entity.SubMenu, error) {
	submenus, err := u.subMenuRepository.GetList(r.MemoIDList)
	return submenus, err
}
func (u subMenuUseCase) BulkCreate(bc requests.BulkCreateSubMenuRequest) error {
	var submenus []entity.SubMenu
	for _, v := range bc.CreateRequests {
		submenu := entity.NewSubMenu(
			entity.SubMenuNameOption(v.Name),
		)
		submenus = append(submenus, *submenu)
	}
	err := u.subMenuRepository.Save(submenus)
	return err
}
