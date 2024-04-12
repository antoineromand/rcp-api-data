package usecase

import (
	"rcp-api-data/internal/entity/domain/data"
	repository "rcp-api-data/internal/repository/data-collector"

	"gorm.io/gorm"
)

type GetCarsUseCase struct {
	DB *gorm.DB
}

type GCResponse struct {
	Success bool
	Code    int
	Data    []data.Car
}

func NewGetCarsUseCase(db *gorm.DB) *GetCarsUseCase {
	return &GetCarsUseCase{
		DB: db,
	}
}

func (e *GetCarsUseCase) GetCars() GCResponse {
	car_repository := repository.NewCarRepository(e.DB)
	cars, err := car_repository.GetAllCars()
	if err != nil {
		return GCResponse{
			Success: false,
			Code:    500,
			Data:    []data.Car{},
		}
	}
	return GCResponse{
		Success: true,
		Code:    200,
		Data:    cars,
	}
}
