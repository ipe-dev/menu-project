package queryservice

import "github.com/ipe-dev/menu_project/domain/dto"

type MemoQueryService interface {
	GetMemoWithAccompanyingData(ID int) (dto.MemoDto, error)
}
