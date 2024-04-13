package usecase

import (
	dto "rcp-api-data/internal/dto/data-collector"
	repository "rcp-api-data/internal/repository/data-collector"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)

type GetCarsUserByIdUseCase struct {
	DB *gorm.DB
}

type GCUResponse struct {
	Success bool
	Message string
	Data    []dto.CarWithBacCount
}

func NewGetCarsUserByIdUseCase(db *gorm.DB) *GetCarsUserByIdUseCase {
	return &GetCarsUserByIdUseCase{
		DB: db,
	}
}

func (g *GetCarsUserByIdUseCase) GetCarsUserById(uuid string) GCUResponse {
	repository := repository.NewCarUserRepository(g.DB)
	_uuid, err := utils.ConvertStringToUUID(uuid)
	if err != nil {
		return GCUResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}
	cars, err := repository.GetCarsWithBacCountByUserUUID(_uuid)

	if err != nil {
		return GCUResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}

	return GCUResponse{
		Success: true,
		Message: "Cars with bac count fetched successfully",
		Data:    cars,
	}

}
