package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Car_User struct {
	gorm.Model
	Id        uint64    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	User_uuid uuid.UUID `json:"user_uuid"`
	Active    bool      `json:"active"`
	CarID     uint      `json:"car"`
	Name      string    `json:"name"`
}

func (Car_User) TableName() string {
	return "car_user"
}

func NewCarUser(user_uuid uuid.UUID, name string) *Car_User {
	return &Car_User{
		User_uuid: user_uuid,
		Active:    true,
		Name:      name,
	}
}
