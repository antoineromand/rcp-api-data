package repository

import (
	entity_account "rcp-api-data/internal/entity/domain/account"

	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
	IAccountRepository
}

type IAccountRepository interface {
	GetAccountByUserUUID(uuid string) (*entity_account.Account, error)
	CreateAccount(account *entity_account.Account) (*entity_account.Account, error)
	UpdateAccount(account *entity_account.Account) error
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
	if result.Error == gorm.ErrRecordNotFound {
		return nil, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return account, nil
}

func (ar *AccountRepository) CreateAccount(account *entity_account.Account) (*entity_account.Account, error) {
	if err := ar.DB.Create(account).Error; err != nil {
		return nil, err
	}
	account, err := ar.GetAccountByUserUUID(account.UserUUID.String())
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (ar *AccountRepository) UpdateAccount(account *entity_account.Account) error {
	var entity *entity_account.Account
	entity, err := ar.GetAccountByUserUUID(account.UserUUID.String())
	if err != nil {
		return err
	}

	if err := ar.DB.Model(&entity).Updates(account).Error; err != nil {
		return err
	}
	return nil
}
