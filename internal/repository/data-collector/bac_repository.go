package repository

import (
	dto "rcp-api-data/internal/dto/data-collector"
	"rcp-api-data/internal/entity/domain/data"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IBacRepository interface {
	CreateBac(*data.Bac) (*data.Bac, error)
	GetBacByID(uint) (data.Bac, error)
	GetAllBacs(uuid.UUID) ([]data.Bac, error)
	UpdateBac(*data.Bac) (*data.Bac, error)
	DeleteBacByID(uint) error
	GetStats(uuid.UUID) (*interface{}, error)
	GetBacsWithLastMeasurementByUserUUID(uuid.UUID, uint) ([]dto.BacWithLastMeasurement, error)
}

type BacRepository struct {
	DB *gorm.DB
	IBacRepository
}

func NewBacRepository(db *gorm.DB) IBacRepository {
	return &BacRepository{
		DB: db,
	}
}

func (br *BacRepository) CreateBac(bac *data.Bac) (*data.Bac, error) {
	if err := br.DB.Create(bac).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func (br *BacRepository) GetBacByID(id uint) (data.Bac, error) {
	bac := data.Bac{}
	result := br.DB.Where("id = ?", id).First(&bac)
	if result.Error == gorm.ErrRecordNotFound {
		return data.Bac{}, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return data.Bac{}, result.Error
	}
	return bac, nil
}

func (br *BacRepository) UpdateBac(bac *data.Bac) (*data.Bac, error) {
	if err := br.DB.Save(bac).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func (br *BacRepository) GetAllBacs(uuid uuid.UUID) ([]data.Bac, error) {
	var bacs []data.Bac
	result := br.DB.Where("car_user_id = ?", uuid).Find(&bacs)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return bacs, nil
}

func (br *BacRepository) DeleteBacByID(id uint) error {
	result := br.DB.Delete(&data.Bac{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (br *BacRepository) GetStats(uuid uuid.UUID) (*interface{}, error) {
	// TODO: Get size from PO and get all measurements to calculate : ex 10/100 * size + ... + ...
	var stats interface{}
	result := br.DB.Raw("SELECT * FROM bacs WHERE car_user_id = ?", uuid).Scan(&stats)
	if result.Error != nil {
		return nil, result.Error
	}
	return &stats, nil
}

func (cur *BacRepository) GetBacsWithLastMeasurementByUserUUID(userUUID uuid.UUID, car_user_id uint) ([]dto.BacWithLastMeasurement, error) {
	var bacs []dto.BacWithLastMeasurement

	result := cur.DB.
		Model(&data.Bac{}).
		Joins("JOIN car_user ON bac.centrale_module_id = car_user.id").
		Joins("JOIN car ON car_user.car_id = car.id").
		Joins("JOIN brand ON car.car_brand_id = brand.id").
		Joins("LEFT JOIN (SELECT bac_id, MAX(created_at) AS last_measurement_date, weight FROM microplastic_measurement GROUP BY bac_id, weight) AS last_measurement ON bac.id = last_measurement.bac_id").
		Where("car_user.user_uuid = ?", userUUID).
		Where("car_user.id = ?", car_user_id).
		Select("bac.id, last_measurement.last_measurement_date AS date, last_measurement.weight AS last_measurement").
		Find(&bacs)

	if result.Error != nil {
		return nil, result.Error
	}

	return bacs, nil
}
