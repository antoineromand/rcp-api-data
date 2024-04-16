package usecase

import (
	repository "rcp-api-data/internal/repository/data-collector"

	"gorm.io/gorm"
)

type ActivateCarUserUseCase struct {
	DB *gorm.DB
}

type ACUResponse struct {
	Success bool
	Message string
}

func NewActivateCarUserUseCase(db *gorm.DB) *ActivateCarUserUseCase {
	return &ActivateCarUserUseCase{
		DB: db,
	}
}

func (d *ActivateCarUserUseCase) ActivateCarUser(id uint) DCUResponse {
	repository := repository.NewCarUserRepository(d.DB)
	car_user, err := repository.GetCarUserByID(id)
	if err != nil {
		return DCUResponse{
			Success: false,
			Message: err.Error(),
		}
	}
	if car_user.Active {
		return DCUResponse{
			Success: false,
			Message: "Car user already activated",
		}
	}
	car_user.Active = true
	_, err = repository.UpdateCarUser(car_user)
	if err != nil {
		return DCUResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	return DCUResponse{
		Success: true,
		Message: "Car user activated successfully",
	}
}
