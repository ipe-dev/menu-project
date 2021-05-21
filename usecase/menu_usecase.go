package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
)

type MenuUseCase interface {
	CreateMenu(c *gin.Context) error
}
type menuUseCase struct {
	menuRepository repository.MenuRepository
}

func NewMenuUseCase(r repository.MenuRepository) MenuUseCase {
	return &menuUseCase{r}
}

func (mu menuUseCase) CreateMenu(c *gin.Context) error {
	var menu entity.Menu
	c.BindJSON(&menu)

	err := mu.menuRepository.Create(&menu)
	return err
}
