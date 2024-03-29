package usecase

import (
	"rcp-api-data/internal/common"
	"rcp-api-data/internal/config/security"
	account_entity "rcp-api-data/internal/entity/domain/account"
	"rcp-api-data/internal/repository"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)

func GetInformationsByUserUuid(db *gorm.DB, token *security.TokenFromCookie) common.Response {
	sugar := utils.GetLogger()
	accountRepository := repository.AccountRepository{DB: db}
	account, err := accountRepository.GetAccountByUserUUID(token.UUID)
	var accountWithCredentials account_entity.AccountWithCredentials
	accountWithCredentials.Username = token.Username
	accountWithCredentials.Email = token.Email
	if err == gorm.ErrRecordNotFound {
		convertedUUID, err := utils.ConvertStringToUUID(token.UUID)
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
		accountWithCredentials.Account = *account
		return common.Response{Code: 200, Error: nil, Data: accountWithCredentials}
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
	accountWithCredentials.Account = *account
	return common.Response{Code: 200, Error: nil, Data: accountWithCredentials}
}
