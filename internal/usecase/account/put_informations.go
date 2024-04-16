package usecase

import (
	"rcp-api-data/internal/common"
	"rcp-api-data/internal/config"
	"rcp-api-data/internal/config/security"
	"rcp-api-data/internal/dto"
	"rcp-api-data/internal/mapper"
	"rcp-api-data/internal/repository"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)

type PutInformationsUsecase struct {
	db  *gorm.DB
	cfg *security.Environment
}

func NewPutInformationsUsecase(db *gorm.DB, cfg *security.Environment) *PutInformationsUsecase {
	return &PutInformationsUsecase{
		db:  db,
		cfg: cfg,
	}
}

func (u *PutInformationsUsecase) PutInformations(_uuid string, account []byte) *common.Response {
	sugar := utils.GetLogger()
	accountRepository := repository.AccountRepository{DB: u.db}
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
	accountMappingResult, err := mapper.AccountMapping(account, &uuid)
	accountEntity := accountMappingResult.Account
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
	if accountEntity.Username != nil || accountEntity.Email != nil || accountMappingResult.Password != nil {
		redPanda := config.NewKafkaService(utils.ConvertEnvStringToArray(u.cfg.RP_BROKERS), "update-credential")
		var userCredentials = dto.NewUserCredentialsDTO(accountEntity.Username, accountMappingResult.Password, accountEntity.Email)
		redPanda.SendMessage(accountEntity.UserUUID.String(), userCredentials)
	}
	return &common.Response{
		Data:  map[string]string{"message": "Account Profile updated successfully"},
		Error: nil,
		Code:  200,
	}
}
