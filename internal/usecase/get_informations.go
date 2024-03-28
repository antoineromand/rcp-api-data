package usecase

import (
	"rcp-api-data/internal/common"
	account_entity "rcp-api-data/internal/entity/domain/account"
	"rcp-api-data/internal/repository"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)

func GetInformationsByUserUuid(db *gorm.DB, uuid string, username string) common.Response {
	sugar := utils.GetLogger()
	accountRepository := repository.AccountRepository{DB: db}
	account, err := accountRepository.GetAccountByUserUUID(uuid)
	if err == gorm.ErrRecordNotFound {
		convertedUUID, err := utils.ConvertStringToUUID(uuid)
		if err != nil {
			sugar.Error("Error while converting string to UUID", err)
			return common.Response{
				Code: 400,
				Error: &common.CustomError{
					Message: "Error while converting string to UUID",
				},
				Data: nil,
			}
		}
		dto := &account_entity.Account{
			UserUUID: convertedUUID,
			Username: &username,
		}
		account, err := accountRepository.CreateAccount(dto)
		if err != nil {
			sugar.Error("Error while creating account profile", err)
			return common.Response{
				Code: 400,
				Error: &common.CustomError{
					Message: "Error while creating account profile",
				},
				Data: nil,
			}
		}
		return common.Response{Code: 200, Error: nil, Data: account}
	}
	if err != nil {
		sugar.Error("Error while getting account profile", err)
		return common.Response{
			Code: 400,
			Error: &common.CustomError{
				Message: "Error while getting account profile",
			},
			Data: nil,
		}
	}
	return common.Response{Code: 200, Error: nil, Data: account}
}
