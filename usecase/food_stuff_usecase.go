package usecase

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/usecase/requests"
)

type FoodStuffUseCase interface {
	Get(requests.GetFoodStuffRequest) (entity.FoodStuff, error)
	GetList(requests.GetFoodStuffListRequest) ([]entity.FoodStuff, error)
	BulkCreate(requests.BulkCreateFoodStuffRequest) error
	ChangeBuyStatus(requests.ChangeFoodStuffStatusRequest) error
}
type foodStuffUseCase struct {
	foodStuffRepository repository.FoodStuffRepository
}

func NewFoodStuffUseCase(r repository.FoodStuffRepository) FoodStuffUseCase {
	return &foodStuffUseCase{r}
}

func (u foodStuffUseCase) Get(r requests.GetFoodStuffRequest) (entity.FoodStuff, error) {
	var foodstuff entity.FoodStuff
	var err error
	foodstuff, err = u.foodStuffRepository.GetByMenuID(r.MenuID)
	return foodstuff, err
}

func (u foodStuffUseCase) GetList(r requests.GetFoodStuffListRequest) ([]entity.FoodStuff, error) {
	foodstuffs, err := u.foodStuffRepository.GetList(r.MenuIDList)
	return foodstuffs, err
}
func (u foodStuffUseCase) BulkCreate(r requests.BulkCreateFoodStuffRequest) error {
	var foodstuffs []entity.FoodStuff
	for _, v := range r.CreateRequests {
		foodstuff := entity.NewFoodStuff(
			entity.FoodStuffMenuNameOption(v.Name),
		)
		foodstuffs = append(foodstuffs, *foodstuff)
	}
	err := u.foodStuffRepository.Save(foodstuffs)
	return err
}

func (u foodStuffUseCase) ChangeBuyStatus(r requests.ChangeFoodStuffStatusRequest) error {
	err := u.foodStuffRepository.ChangeBuyStatus(r.ID, r.Status)
	return err
}
