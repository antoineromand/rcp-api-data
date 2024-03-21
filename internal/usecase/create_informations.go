package usecase

import (
	"errors"
	entity "rcp-api-data/internal/entity/domain/data"
	"rcp-api-data/internal/dto"
	"rcp-api-data/internal/entity/domain/data/service"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)

func CreateInformations(db *gorm.DB, uuid string, account dto.AccountDTO) {
	sugar := utils.GetLogger()
	
}