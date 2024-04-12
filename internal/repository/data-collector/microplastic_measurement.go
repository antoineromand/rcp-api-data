package repository

import (
	"rcp-api-data/internal/entity/domain/data"

	"gorm.io/gorm"
)

type IMicroplasticMeasurement interface {
	CreateMicroplasticMeasurement(*data.MicroplasticMeasurement) (*data.MicroplasticMeasurement, error)
	GetMicroplasticMeasurementByID(uint) (*data.MicroplasticMeasurement, error)
}

type MicroplasticMeasurementRepository struct {
	DB *gorm.DB
	IMicroplasticMeasurement
}

func NewMicroplasticMeasurementRepository(db *gorm.DB) IMicroplasticMeasurement {
	return &MicroplasticMeasurementRepository{
		DB: db,
	}
}

func (mmr *MicroplasticMeasurementRepository) CreateMicroplasticMeasurement(microplastic_measurement *data.MicroplasticMeasurement) (*data.MicroplasticMeasurement, error) {
	if err := mmr.DB.Create(microplastic_measurement).Error; err != nil {
		return nil, err
	}
	return nil, nil
}