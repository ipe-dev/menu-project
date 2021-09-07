package persistence

import (
	"github.com/ipe-dev/menu_project/domain/dto"
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/queryservice"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/errors"
	"github.com/ipe-dev/menu_project/infrastructure/database"
)

type memoPersistence struct{}
type memoQueryService struct{}

func NewMemoPersistence() repository.MemoRepository {
	return &memoPersistence{}
}

func NewMemoQueryService() queryservice.MemoQueryService {
	return &memoQueryService{}
}
func (p memoPersistence) Create(memo entity.Memo) error {
	tx := database.Db.Begin()
	if err := tx.Model(&memo).Create(memo).Error; err != nil {
		tx.Rollback()
		return errors.NewInfraError(err, memo)
	}
	tx.Commit()
	return nil
}
func (p memoPersistence) Update(memo entity.Memo) error {
	tx := database.Db.Begin()

	if err := tx.Model(&memo).Updates(memo).Error; err != nil {
		tx.Rollback()
		return errors.NewInfraError(err, memo)
	}
	tx.Commit()
	return nil
}

func (p memoPersistence) GetByID(id int) (entity.Memo, error) {
	var memo entity.Memo
	Db := database.Db
	err := Db.Table("memos").Where("id = ?", id).First(&memo).Error
	if err != nil {
		return memo, errors.NewInfraError(err, id)
	}
	return memo, nil
}

func (p memoPersistence) GetList(userID int) ([]entity.Memo, error) {
	var memo []entity.Memo
	Db := database.Db
	err := Db.Where("user_id = ?", userID).Find(&memo).Error
	if err != nil {
		return memo, errors.NewInfraError(err, userID)
	}
	return memo, nil
}
func (q memoQueryService) GetMemoWithAccompanyingData(ID int) (dto.MemoDto, error) {
	Db := database.Db
	var memoDto dto.MemoDto

	err := Db.Table("memos").Where("id = ?", ID).
		Joins("left join menus on memos.id = menus.memo_id").
		Joins("left join sub_menus on memos.id = sub_menus.memo_id").
		Joins("left join food_stuffs on memos.id = food_stuffs.memo_id").
		Find(memoDto).
		Error
	if err != nil {
		return memoDto, errors.NewInfraError(err, ID)
	}
	return memoDto, nil
}
