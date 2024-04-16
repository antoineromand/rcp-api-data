package usecase

import (
	"rcp-api-data/internal/entity/domain/data"
	mapper "rcp-api-data/internal/mapper/data-collector"
	repository "rcp-api-data/internal/repository/data-collector"

	"gorm.io/gorm"
)

type InsertCarUser struct {
	DB *gorm.DB
}

type ICResponse struct {
	Success bool
	Code    int
	Message string
}

func NewInsertCarUserUseCase(db *gorm.DB) *InsertCarUser {
	return &InsertCarUser{
		DB: db,
	}
}

func (e *InsertCarUser) InsertCarUser(uuid string, bytes []byte) ICResponse {
	car_user_repository := repository.NewCarUserRepository(e.DB)
	dto, err := mapper.CarUserMapping(bytes, uuid)
	if err != nil {
		return ICResponse{
			Success: false,
			Code:    500,
			Message: "Error mapping car user",
		}
	}
	_, err = car_user_repository.CreateCarUser(
		&data.Car_User{
			User_uuid: dto.UserUUID,
			CarID:     dto.CarID,
			Active:    true,
		},
	)
	if err != nil {
		return ICResponse{
			Success: false,
			Code:    500,
			Message: "Error creating car user",
		}
	}
	return ICResponse{
		Success: true,
		Code:    200,
		Message: "Car user created",
	}
}
