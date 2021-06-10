package usecase

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
)

type SubFoodStuffUseCase interface {
	GetSubFoodStuff(GetSubFoodStuffRequest) (entity.SubFoodStuff, error)
	GetSubFoodStuffList(GetSubFoodStuffListRequest) ([]entity.SubFoodStuff, error)
	BulkCreateSubFoodStuff(BulkCreateSubFoodStuffRequest) ([]entity.SubFoodStuff, error)
	BulkUpdateSubFoodStuff(BulkUpdateSubFoodStuffRequest) ([]entity.SubFoodStuff, error)
	ChangeBuyStatus(ChangeSubBuyStatusRequest) error
}
type subFoodStuffUseCase struct {
	subFoodStuffRepository repository.SubFoodStuffRepository
}

func NewSubFoodStuffUseCase(r repository.SubFoodStuffRepository) SubFoodStuffUseCase {
	return &subFoodStuffUseCase{r}
}

type GetSubFoodStuffRequest struct {
	SubMenuID int `json:"menu_id"`
}
type GetSubFoodStuffListRequest struct {
	SubMenuIDList []int `json:"menu_id_list"`
}
type CreateSubFoodStuffRequest struct {
	Name      string `json:"name"`
	SubMenuID int    `json:"menu_id"`
}
type BulkCreateSubFoodStuffRequest struct {
	CreateRequests []CreateSubFoodStuffRequest
}
type UpdateSubFoodStuffRequest struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	SubMenuID int    `json:"menu_id"`
}
type BulkUpdateSubFoodStuffRequest struct {
	UpdateRequests []UpdateSubFoodStuffRequest
}
type ChangeSubBuyStatusRequest struct {
	ID     int `json:"id"`
	Status int `json:"status"`
}

func (u subFoodStuffUseCase) GetSubFoodStuff(r GetSubFoodStuffRequest) (entity.SubFoodStuff, error) {
	var foodstuff entity.SubFoodStuff
	var err error
	foodstuff, err = u.subFoodStuffRepository.GetBySubMenuID(r.SubMenuID)
	return foodstuff, err
}

func (u subFoodStuffUseCase) GetSubFoodStuffList(r GetSubFoodStuffListRequest) ([]entity.SubFoodStuff, error) {
	foodstuffs, err := u.subFoodStuffRepository.GetList(r.SubMenuIDList)
	return foodstuffs, err
}
func (u subFoodStuffUseCase) BulkCreateSubFoodStuff(r BulkCreateSubFoodStuffRequest) ([]entity.SubFoodStuff, error) {
	var foodstuffs []entity.SubFoodStuff
	for _, v := range r.CreateRequests {
		foodstuff := entity.SubFoodStuff{
			Name:      v.Name,
			SubMenuID: v.SubMenuID,
		}
		foodstuffs = append(foodstuffs, foodstuff)
	}
	foodstuffs, err := u.subFoodStuffRepository.BulkCreate(foodstuffs)
	return foodstuffs, err
}
func (u subFoodStuffUseCase) BulkUpdateSubFoodStuff(r BulkUpdateSubFoodStuffRequest) ([]entity.SubFoodStuff, error) {
	var foodstuffs []entity.SubFoodStuff
	for _, v := range r.UpdateRequests {
		foodstuff := entity.SubFoodStuff{
			ID:        v.ID,
			Name:      v.Name,
			SubMenuID: v.SubMenuID,
		}
		foodstuffs = append(foodstuffs, foodstuff)
	}
	foodstuffs, err := u.subFoodStuffRepository.BulkUpdate(foodstuffs)
	return foodstuffs, err
}
func (u subFoodStuffUseCase) ChangeBuyStatus(r ChangeSubBuyStatusRequest) error {
	err := u.subFoodStuffRepository.ChangeBuyStatus(r.ID, r.Status)
	return err
}
