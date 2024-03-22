package usecase

import (
	"rcp-api-data/internal/common"
	"rcp-api-data/internal/dto"
	"rcp-api-data/internal/repository"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)

func CreateInformations(db *gorm.DB, uuid string, account dto.AccountDTO) *common.Response {
	sugar := utils.GetLogger()
	accountRepository := repository.AccountRepository{DB: db}
	err := accountRepository.CreateAccount(&account, uuid)
	if err != nil {
		sugar.Error("Error while creating account profile", err)
		return &common.Response{
			Data: nil,
			Error: &common.CustomError{
				Message: "Error while creating account profile",
			},
			Code: 400,
		}
	}
	return &common.Response{
		Data: map[string]string{"message": "Account Profile created successfully"},
		Error: nil,
		Code: 200,
	}
}