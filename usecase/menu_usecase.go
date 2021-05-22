package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
)

type MenuUseCase interface {
	CreateMenu(CreateMenuRequest) error
	GetMenu(c *gin.Context) (*entity.Menu, error)
}
type menuUseCase struct {
	menuRepository repository.MenuRepository
}
type CreateMenuRequest struct {
	Name   string `json:"name"`
	Date   string `json:"date"`
	Kind   int    `json:"kind"`
	URL    string `json:"url"`
	UserID int    `json:"user_id"`
}
type GetMenuListRequest struct {
	StartDate int `json:"start_date"`
	UserID    int `json:"user_id"`
}
type GetMenuRequest struct {
	ID int `json:"id"`
}
type GetListRequest struct {
	Date int `json:"date"`
}

func newMenu()

func NewMenuUseCase(r repository.MenuRepository) MenuUseCase {
	return &menuUseCase{r}
}

func (mu menuUseCase) CreateMenu(mr CreateMenuRequest) error {
	err := mu.menuRepository.Create(mr)
	return err
}

func (mu menuUseCase) GetMenu(c *gin.Context) (*entity.Menu, error) {

	return
}
