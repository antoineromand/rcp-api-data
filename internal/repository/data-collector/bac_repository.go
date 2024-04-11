package repository

import (
	dto "rcp-api-data/internal/dto/data-collector"
	"rcp-api-data/internal/entity/domain/data"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IBacRepository interface {
	CreateBac(*data.Bac) (*data.Bac, error)
	GetBacByID(uint) (*data.Bac, error)
	GetAllBacs(uuid.UUID) ([]data.Bac, error)
	UpdateBac(*data.Bac) (*data.Bac, error)
	DeleteBacByID(uint) error
	GetStats(uuid.UUID) (*interface{}, error)
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

func (br *BacRepository) GetBacByID(id uint) (*data.Bac, error) {
	bac := &data.Bac{}
	result := br.DB.Where("id = ?", id).First(bac)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return nil, result.Error
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
	var stats interface{}
	result := br.DB.Raw("SELECT * FROM bacs WHERE car_user_id = ?", uuid).Scan(&stats)
	if result.Error != nil {
		return nil, result.Error
	}
	return &stats, nil
}

func (cur *BacRepository) GetBacsWithLastMeasurementByUserUUID(userUUID uuid.UUID) ([]dto.BacWithLastMeasurement, error) {
    var bacs []dto.BacWithLastMeasurement

    result := cur.DB.
        Model(&data.Bac{}).
        Joins("JOIN car_user ON bac.centrale_module_id = car_user.id").
        Joins("JOIN car ON car_user.car_id = car.id").
        Joins("JOIN brand ON car.car_brand_id = brand.id").
        Where("car_user.user_uuid = ?", userUUID).
        Select("bac.id, bac.name, (SELECT mm.weight FROM microplastic_measurement mm WHERE mm.bac_id = bac.id ORDER BY mm.created_at DESC LIMIT 1) AS last_measurement").
        Find(&bacs)

    if result.Error != nil {
        return nil, result.Error
    }

    return bacs, nil
}