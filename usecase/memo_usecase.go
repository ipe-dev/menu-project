package usecase

import (
	"github.com/ipe-dev/menu_project/domain/dto"
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/entity/value"
	"github.com/ipe-dev/menu_project/domain/queryservice"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/usecase/requests"
)

type MemoUseCase interface {
	GetList(requests.GetMemoListRequest) ([]entity.Memo, error)
	Create(requests.CreateMemoRequest) error
	Update(requests.UpdateMemoRequest) error
	Get(requests.GetMemoRequest) (dto.Memo, error)
}
type memoUseCase struct {
	memoRepository   repository.MemoRepository
	memoQueryService queryservice.MemoQueryService
}

func NewMemoUseCase(r repository.MemoRepository, q queryservice.MemoQueryService) MemoUseCase {
	return memoUseCase{
		memoRepository:   r,
		memoQueryService: q,
	}
}

func (u memoUseCase) GetList(r requests.GetMemoListRequest) ([]entity.Memo, error) {
	memos, err := u.memoRepository.GetList(r.UserID)
	return memos, err
}

func (u memoUseCase) Create(r requests.CreateMemoRequest) error {
	memo, err := entity.NewMemo(
		entity.MemoTitleOption(value.NewTitle(r.StartDate, r.EndDate)),
		entity.MemoStartDateOption(r.StartDate),
		entity.MemoEndDateOption(r.EndDate),
	)
	if err != nil {
		return err
	}
	err = u.memoRepository.Create(*memo)
	return err
}
func (u memoUseCase) Update(r requests.UpdateMemoRequest) error {
	memo, err := entity.NewMemo(
		entity.MemoIDOption(r.ID),
		entity.MemoTitleOption(value.NewTitle(r.StartDate, r.EndDate)),
		entity.MemoStartDateOption(r.StartDate),
		entity.MemoEndDateOption(r.EndDate),
	)
	if err != nil {
		return err
	}
	err = u.memoRepository.Update(*memo)
	return nil
}
func (u memoUseCase) Get(r requests.GetMemoRequest) (dto.Memo, error) {
	memo, err := u.memoQueryService.GetMemoWithAccompanyingData(r.ID, r.UserID)
	if err == nil {
		memo.ConvertDateTimestamp()
	}
	return memo, err
}
