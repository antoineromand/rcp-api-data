package repository

import (
	"rcp-api-data/internal/entity/domain/data"

	"gorm.io/gorm"
)

type ICentraleModuleRepository interface {
	CreateCentraleModule(*data.CentraleModule) (*uint, error)
	GetCentraleModuleByModuleSSID(ssid string) (*data.CentraleModule, error)
	GetCentraleModuleIdBySSID(ssid string) (*uint, error)
	GetAllCentraleModuleByCarUserUUID(*uint) ([]data.CentraleModule, error)
	GetSSIDByCarUserId(id uint) (*string, error)
}

type CentraleModuleRepository struct {
	DB *gorm.DB
	ICentraleModuleRepository
}

func NewModuleRepository(db *gorm.DB) ICentraleModuleRepository {
	return &CentraleModuleRepository{
		DB: db,
	}
}

// TODO: Create a method that gonna insert a module into the database

func (cmr *CentraleModuleRepository) CreateCentraleModule(centrale_module *data.CentraleModule) (*uint, error) {
	if err := cmr.DB.Create(centrale_module).Error; err != nil {
		return nil, err
	}
	return &centrale_module.ID, nil
}

func (cmr *CentraleModuleRepository) GetCentraleModuleByModuleSSID(ssid string) (*data.CentraleModule, error) {
	module := &data.CentraleModule{}
	result := cmr.DB.Where("ss_id = ?", ssid).First(module)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return module, nil
}

func (cmr *CentraleModuleRepository) GetSSIDByCarUserId(id uint) (*string, error) {
	module := &data.CentraleModule{}
	result := cmr.DB.Where("car_user_id = ?", id).First(module)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &module.SSID, nil
}

func (cmr *CentraleModuleRepository) GetAllCentraleModuleByCarUserUUID(car_user_id *uint) ([]data.CentraleModule, error) {
	var modules []data.CentraleModule
	result := cmr.DB.Where("car_user_id = ?", car_user_id).Find(&modules)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return modules, nil
}

func (cmr *CentraleModuleRepository) GetCentraleModuleIdBySSID(ssid string) (*uint, error) {
	module := &data.CentraleModule{}
	result := cmr.DB.Where("ss_id = ?", ssid).First(module)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &module.ID, nil
}
