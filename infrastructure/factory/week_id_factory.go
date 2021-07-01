package factory

import (
	"github.com/ipe-dev/menu_project/infrastructure/database"
)

type WeekIDFactory interface {
	NewWeekID(UserID int) (int, error)
	IncrementWeekID(UserID int) error
}

type weekIDFactory struct{}

func NewWeekIDFactory() WeekIDFactory {
	return weekIDFactory{}
}

func (w weekIDFactory) NewWeekID(UserID int) (int, error) {
	var WeekID struct {
		WeekID int
	}
	tx := database.Db.Begin()
	err := tx.Table("week_ids").Where("user_id = ?", UserID).Find(&WeekID).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if WeekID.WeekID == 0 {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()

	return WeekID.WeekID, err
}
func (w weekIDFactory) IncrementWeekID(UserID int) error {
	tx := database.Db.Begin()
	var WeekIDs []int
	err := tx.Table("week_ids").Where("user_id = ?", UserID).Pluck("week_id", &WeekIDs).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if WeekIDs[0] == 0 {
		tx.Rollback()
		return err
	}
	err = tx.Table("week_ids").Where("user_id = ?", UserID).Update("week_id", WeekIDs[0]+1).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
