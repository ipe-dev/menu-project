package service

import (
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/errors"
)

type MemoService interface {
	CheckMemoExists(ID int) error
}
type memoService struct {
	MemoRepository repository.MemoRepository
}

func NewMemoService(r repository.MemoRepository) MemoService {
	return memoService{r}
}

func (s memoService) CheckMemoExists(ID int) error {
	memo, err := s.MemoRepository.GetByID(ID)
	if err != nil {
		return errors.NewInfraError(err, ID)
	}

	if memo.ID == 0 {
		return errors.NewExistError("メモが存在しません", ID)
	}
	return nil
}
