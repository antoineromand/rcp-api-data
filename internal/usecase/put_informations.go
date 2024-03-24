package usecase

import (
	"rcp-api-data/internal/common"
	"rcp-api-data/internal/mapper"
	"rcp-api-data/internal/repository"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)

func PutInformations(db *gorm.DB, _uuid string, account []byte) *common.Response {
	sugar := utils.GetLogger()
	accountRepository := repository.AccountRepository{DB: db}
	uuid, err := utils.ConvertStringToUUID(_uuid)
	if err != nil {
		sugar.Error("Error while converting string to UUID", err)
		return &common.Response{
			Data: nil,
			Error: &common.CustomError{
				Message: "Error while converting string to UUID",
			},
			Code: 400,
		}
	}
	accountEntity, err := mapper.AccountMapping(account, &uuid)
	if err != nil {
		sugar.Error("Error while mapping account", err)
		return &common.Response{
			Data: nil,
			Error: &common.CustomError{
				Message: "Error while mapping account",
			},
			Code: 400,
		}
	}
	err = accountRepository.UpdateAccount(&accountEntity)
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
