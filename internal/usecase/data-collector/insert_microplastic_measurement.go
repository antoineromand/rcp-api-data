package usecase

import (
	"errors"
	entity "rcp-api-data/internal/entity/domain/data"
	"rcp-api-data/internal/entity/domain/data/service"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)

func InsertMicroplasticMeasurement(db *gorm.DB, measurement *entity.MicroplasticMeasurement) error {
	sugar := utils.GetLogger()
	if measurement == nil {
		sugar.Error("Measurement is nil, please provide a valid measurement")
		return errors.New("measurement is nil")
	}
	if db == nil {
		sugar.Error("Measurement is nil, please provide a valid measurement")
		return errors.New("database is nil")
	}
	// Get bac by id
	bacService := service.NewBacService(db)
	bac, err := bacService.GetBacByID(measurement.BacID)
	if err != nil {
		sugar.Error("Error while getting bac by id", err)
		return errors.New("error while getting bac by id")
	}
	if bac == nil {
		sugar.Error("Bac not found")
		return errors.New("bac not found")
	}
	insertMicroplasticsMeasurement := db.Model(&entity.MicroplasticMeasurement{}).Create(measurement)
	if insertMicroplasticsMeasurement.Error != nil {
		sugar.Error("Error while inserting microplastic measurement", insertMicroplasticsMeasurement.Error)
		return errors.New("error while inserting microplastic measurement")
	}
	return nil
}