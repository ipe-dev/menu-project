package usecase

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
)

type MemoUseCase interface {
	GetList(GetMemoListRequest) ([]entity.Memo, error)
	Create(CreateMemoRequest) (entity.Memo, error)
	Update(UpdateMemoRequest) (entity.Memo, error)
	Get(GetMemoRequest) (entity.Memo, error)
}
type memoUseCase struct {
	memoRepository repository.MemoRepository
}

func NewMemoUseCase(r repository.MemoRepository) MemoUseCase {
	return memoUseCase{
		memoRepository: r,
	}
}

type CreateMemoRequest struct {
	Title     string `json:"title"`
	StartDate int64  `json:"start_date"`
	EndDate   int64  `json:"end_date"`
}
type UpdateMemoRequest struct {
	ID        int    `json:"id" validate:"required"`
	Title     string `json:"title"`
	StartDate int64  `json:"start_date"`
	EndDate   int64  `json:"end_date"`
}
type GetMemoListRequest struct {
	UserID int `json:"user_id"`
}
type GetMemoRequest struct {
	ID     int `json:"id" validate:"required"`
	UserID int
}

func (u memoUseCase) GetList(r GetMemoListRequest) ([]entity.Memo, error) {
	memos, err := u.memoRepository.GetList(r.UserID)
	if err != nil {
		return memos, err
	}
	return memos, nil
}

func (u memoUseCase) Create(r CreateMemoRequest) (entity.Memo, error) {

	memo := entity.NewMemo(
		entity.MemoTitleOption(r.Title),
		entity.MemoStartDateOption(r.StartDate),
		entity.MemoEndDateOption(r.EndDate),
	)

	memoData, err := u.memoRepository.Create(*memo)

	return memoData, err
}
func (u memoUseCase) Update(r UpdateMemoRequest) (entity.Memo, error) {
	memo := entity.NewMemo(
		entity.MemoIDOption(r.ID),
		entity.MemoTitleOption(r.Title),
		entity.MemoStartDateOption(r.StartDate),
		entity.MemoEndDateOption(r.EndDate),
	)
	memoData, err := u.memoRepository.Update(*memo)

	return memoData, err
}
func (u memoUseCase) Get(r GetMemoRequest) (entity.Memo, error) {
	var memo entity.Memo
	var err error
	if r.ID != 0 {
		memo, err = u.memoRepository.GetByID(r.ID)
	}
	return memo, err
}
