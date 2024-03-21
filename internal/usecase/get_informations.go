package usecase

import (
	entity_account "rcp-api-data/internal/entity/domain/account"
	"rcp-api-data/internal/repository"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)


func Get_informations_by_user_uuid(db *gorm.DB, uuid string) (*entity_account.Account, error) {
	sugar := utils.GetLogger()
	accountRepository := repository.AccountRepository{DB: db}
	account, err := accountRepository.GetAccountByUserUUID(uuid)
	if err != nil {
		sugar.Error("Error while getting account by user uuid", err)
		return nil, err
	}
	if account == nil {
		sugar.Error("Account not found")
		return nil, nil
	}
	return account, nil
}