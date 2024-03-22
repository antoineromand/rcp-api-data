package usecase

import (
	"rcp-api-data/internal/common"
	"rcp-api-data/internal/dto"
	"rcp-api-data/internal/repository"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)

func PutInformations(db *gorm.DB, _uuid string, account dto.AccountDTO) *common.Response {
	sugar := utils.GetLogger()
	accountRepository := repository.AccountRepository{DB: db}
	err := accountRepository.UpdateAccount(&account, _uuid)
	if err != nil {
		sugar.Error("Error while updating account profile", err)
		return &common.Response{
			Data: nil,
			Error: &common.CustomError{
				Message: "Error while updating account profile",
			},
			Code: 400,
		}
	}
	return &common.Response{
		Data: map[string]string{"message": "Account Profile updated successfully"},
		Error: nil,
		Code: 200,
	}
}
