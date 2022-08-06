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
func (q memoQueryService) GetMemoWithAccompanyingData(ID int, UserID int) (dto.Memo, error) {
	Db := database.Db
	var memoDto dto.Memo

	// メモ取得
	err := Db.Model(&memoDto).
		Where("memos.id = ?", ID).
		Where("user_id = ?", UserID).
		Order("memos.created_at desc").
		First(&memoDto).
		Error
	if err != nil {
		return memoDto, errors.NewInfraError(err, ID)
	}
	// メニュー取得
	if err = Db.Model(&memoDto).Association("Menus").Find(&memoDto.Menus); err != nil {
		return memoDto, errors.NewInfraError(err, memoDto)
	}
	// 副菜取得
	if err = Db.Model(&memoDto).Association("SubMenus").Find(&memoDto.SubMenus); err != nil {
		return memoDto, errors.NewInfraError(err, memoDto)
	}
	// 食材取得
	if err = Db.Model(&memoDto).Association("FoodStuffs").Find(&memoDto.FoodStuffs); err != nil {
		return memoDto, errors.NewInfraError(err, memoDto)
	}
	return memoDto, nil
}
