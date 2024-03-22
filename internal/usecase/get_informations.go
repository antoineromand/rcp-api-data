package usecase

import (
	"rcp-api-data/internal/common"
	"rcp-api-data/internal/repository"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)


func GetInformationsByUserUuid(db *gorm.DB, uuid string) common.Response {
	sugar := utils.GetLogger()
	accountRepository := repository.AccountRepository{DB: db}
	account, err := accountRepository.GetAccountByUserUUID(uuid)
	if err != nil {
		sugar.Error("Error while getting account by user uuid, account not found", err)
		return common.Response{Code: 500, Error: &common.CustomError{
			Message: "Error while getting account by user uuid, account not found",
		}, Data: nil}
	}
	return common.Response{Code: 200, Error: nil, Data: account}
}