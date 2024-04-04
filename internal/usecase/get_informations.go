package usecase

import (
	"rcp-api-data/internal/common"
	"rcp-api-data/internal/config/security"
	account_entity "rcp-api-data/internal/entity/domain/account"
	"rcp-api-data/internal/repository"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)

type GetInformationsUseCase struct {
	db *gorm.DB
}

func NewGetInformationsUseCase(db *gorm.DB) *GetInformationsUseCase {
	return &GetInformationsUseCase{
		db: db,
	}
}

func (e *GetInformationsUseCase) GetInformationsByUserUuid(token *security.TokenFromCookie) common.Response {
	sugar := utils.GetLogger()
	accountRepository := repository.AccountRepository{DB: e.db}
	account, err := accountRepository.GetAccountByUserUUID(token.UUID)
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
			Username: token.Username,
			Email:    token.Email,
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
