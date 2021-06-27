package usecase

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
)

type FoodStuffUseCase interface {
	Get(GetFoodStuffRequest) (entity.FoodStuff, error)
	GetList(GetFoodStuffListRequest) ([]entity.FoodStuff, error)
	BulkCreate(BulkCreateFoodStuffRequest) ([]entity.FoodStuff, error)
	BulkUpdate(BulkUpdateFoodStuffRequest) ([]entity.FoodStuff, error)
	ChangeBuyStatus(ChangeFoodStuffStatusRequest) error
}
type foodStuffUseCase struct {
	foodStuffRepository repository.FoodStuffRepository
}

func NewFoodStuffUseCase(r repository.FoodStuffRepository) FoodStuffUseCase {
	return &foodStuffUseCase{r}
}

type GetFoodStuffRequest struct {
	MenuID int `json:"menu_id" validate:"required"`
}
type GetFoodStuffListRequest struct {
	MenuIDList []int `json:"menu_id_list"`
}
type CreateFoodStuffRequest struct {
	Name   string `json:"name" validate:"required"`
	MenuID int    `json:"menu_id" validate:"required"`
}
type BulkCreateFoodStuffRequest struct {
	CreateRequests []CreateFoodStuffRequest
}
type UpdateFoodStuffRequest struct {
	ID     int    `json:"id" validate:"required"`
	Name   string `json:"name"`
	MenuID int    `json:"menu_id" validate:"required"`
}
type BulkUpdateFoodStuffRequest struct {
	UpdateRequests []UpdateFoodStuffRequest
}
type ChangeFoodStuffStatusRequest struct {
	ID     int
	Status int
}

func (u foodStuffUseCase) Get(r GetFoodStuffRequest) (entity.FoodStuff, error) {
	var foodstuff entity.FoodStuff
	var err error
	foodstuff, err = u.foodStuffRepository.GetByMenuID(r.MenuID)
	return foodstuff, err
}

func (u foodStuffUseCase) GetList(r GetFoodStuffListRequest) ([]entity.FoodStuff, error) {
	foodstuffs, err := u.foodStuffRepository.GetList(r.MenuIDList)
	return foodstuffs, err
}
func (u foodStuffUseCase) BulkCreate(r BulkCreateFoodStuffRequest) ([]entity.FoodStuff, error) {
	var foodstuffs []entity.FoodStuff
	for _, v := range r.CreateRequests {
		foodstuff := entity.NewFoodStuff(
			entity.FoodStuffMenuNameOption(v.Name),
			entity.FoodStuffMenuIDOption(v.MenuID),
		)
		foodstuffs = append(foodstuffs, *foodstuff)
	}
	foodstuffs, err := u.foodStuffRepository.BulkCreate(foodstuffs)
	return foodstuffs, err
}
func (u foodStuffUseCase) BulkUpdate(r BulkUpdateFoodStuffRequest) ([]entity.FoodStuff, error) {
	var foodstuffs []entity.FoodStuff
	for _, v := range r.UpdateRequests {
		foodstuff := entity.NewFoodStuff(
			entity.FoodStuffIDOption(v.ID),
			entity.FoodStuffMenuNameOption(v.Name),
			entity.FoodStuffMenuIDOption(v.MenuID),
		)
		foodstuffs = append(foodstuffs, *foodstuff)
	}
	foodstuffs, err := u.foodStuffRepository.BulkUpdate(foodstuffs)
	return foodstuffs, err
}

func (u foodStuffUseCase) ChangeBuyStatus(r ChangeFoodStuffStatusRequest) error {
	err := u.foodStuffRepository.ChangeBuyStatus(r.ID, r.Status)
	return err
}
