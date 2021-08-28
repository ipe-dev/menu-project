package persistence

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/errors"
	"github.com/ipe-dev/menu_project/infrastructure/database"
)

type memoPersistence struct{}

func NewMemoPersistence() repository.MemoRepository {
	return &memoPersistence{}
}
func (p memoPersistence) Create(memo entity.Memo) (entity.Memo, error) {
	tx := database.Db.Begin()
	err := tx.Model(&memo).Create(memo).Error
	if err != nil {
		tx.Rollback()
		return memo, errors.NewInfraError(err, memo)
	}
	tx.Commit()
	return memo, nil
}
func (p memoPersistence) Update(memo entity.Memo) (entity.Memo, error) {
	tx := database.Db.Begin()
	var res entity.Memo
	var err error

	err = tx.Model(&memo).Updates(memo).Error
	if err != nil {
		tx.Rollback()
		return res, errors.NewInfraError(err, memo)
	}
	tx.Commit()
	return res, nil
}

func (p memoPersistence) GetByID(id int, userID int) (entity.Memo, error) {
	var memo entity.Memo
	Db := database.Db
	err := Db.Table("memos").Where("id = ?", id).Where("user_id = ?", userID).First(&memo).Error
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
