package usecase

import (
	repository "rcp-api-data/internal/repository/data-collector"

	"gorm.io/gorm"
)

type DesactivateCarUserUseCase struct {
	DB *gorm.DB
}

type DCUResponse struct {
	Success bool
	Message string
}

func NewDesactivateCarUserUseCase(db *gorm.DB) *DesactivateCarUserUseCase {
	return &DesactivateCarUserUseCase{
		DB: db,
	}
}

func (d *DesactivateCarUserUseCase) DesactivateCarUser(id uint, user_uuid string) DCUResponse {
	repository := repository.NewCarUserRepository(d.DB)
	car_user, err := repository.GetCarUserByID(id)
	if err != nil {
		return DCUResponse{
			Success: false,
			Message: err.Error(),
		}
	}
	if car_user.User_uuid.String() != user_uuid {
		return DCUResponse{
			Success: false,
			Message: "You are not allowed to desactivate this car user",
		}
	}
	if !car_user.Active {
		return DCUResponse{
			Success: false,
			Message: "Car user already desactivated",
		}
	}

	car_user.Active = false
	_, err = repository.UpdateCarUser(car_user)
	if err != nil {
		return DCUResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	return DCUResponse{
		Success: true,
		Message: "Car user desactivated successfully",
	}
}
