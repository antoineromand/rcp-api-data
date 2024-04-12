package dto

import "github.com/google/uuid"

type CarUserDTO struct {
	UserUUID uuid.UUID
	CarID    uint `json:"car_id"`
}
