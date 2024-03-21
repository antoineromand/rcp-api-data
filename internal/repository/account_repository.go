package repository

import (
	"rcp-api-data/internal/dto"
	entity_account "rcp-api-data/internal/entity/domain/account"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
	IAccountRepository
}

type IAccountRepository interface {
	GetAccountByUserUUID(uuid string) (*entity_account.Account, error)
	CreateAccount(account *dto.AccountDTO, _uuid string) error
	UpdateAccount(account *dto.AccountDTO, _uuid string) error
	DeleteAccount(_uuid string) error
}

func NewAccountRepository(db *gorm.DB) IAccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

func (ar *AccountRepository) GetAccountByUserUUID(uuid string) (*entity_account.Account, error) {
	account := &entity_account.Account{}
	result := ar.DB.Where("user_uuid = ?", uuid).First(account)
	if result.Error != nil {
		return nil, result.Error
	}
	return account, nil
}

func (ar *AccountRepository) CreateAccount(account *dto.AccountDTO, _uuid string) error {
	
	userUUID, err := utils.ConvertStringToUUID(_uuid)
	if err != nil {
		return err
	}

	dto := &entity_account.Account{
		UserUUID: userUUID,
		Username: *account.Username,        
		Email: *account.Email,      
		ActivityMessage: *account.ActivityMessage,
		Address: *account.Address,      
		City: *account.City,      
		Country: *account.Country,    
		PostalCode: *account.PostalCode,   
		PhoneNumber: *account.PhoneNumber,     
		FirstName: *account.FirstName,      
		LastName: *account.LastName,      
		IsNew: *account.IsNew, 
	}
	if err := ar.DB.Create(dto).Error; err != nil {
        return err
    }
	return nil
}