package usecase

import (
	dto "rcp-api-data/internal/dto/data-collector"
	repository "rcp-api-data/internal/repository/data-collector"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)

type GetDataCollectorUseCase struct {
	DB *gorm.DB
}

type GDBResponse struct {
	Success bool
	Message string
	SSID    *string
	Data    []dto.BacWithLastMeasurement
}

func NewGetDataCollectorUseCase(db *gorm.DB) *GetDataCollectorUseCase {
	return &GetDataCollectorUseCase{
		DB: db,
	}
}

func (dc *GetDataCollectorUseCase) GetDataByBacs(token string, car_user_id uint) GDBResponse {
	uuid, err := utils.ConvertStringToUUID(token)
	if err != nil {
		return GDBResponse{
			Success: false,
			Message: "Cannot convert token to UUID",
			Data:    nil,
		}
	}
	result, err := repository.NewBacRepository(dc.DB).GetBacsWithLastMeasurementByUserUUID(uuid, car_user_id)
	if err != nil {
		return GDBResponse{
			Success: false,
			SSID:    nil,
			Message: "Cannot get data",
			Data:    nil,
		}
	}
	SSID, err := repository.NewModuleRepository(dc.DB).GetSSIDByCarUserId(car_user_id)
	if err != nil {
		return GDBResponse{
			Success: false,
			SSID:    nil,
			Message: "Cannot get SSID",
			Data:    nil,
		}
	}

	return GDBResponse{
		Success: true,
		SSID:    SSID,
		Message: "Success",
		Data:    result,
	}
}
