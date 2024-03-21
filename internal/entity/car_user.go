package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Car_User struct {
	gorm.Model
	Id uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	CarID uint `json:"car_id"`
	User_uuid uuid.UUID `json:"user_uuid"`
	Name string `json:"name"`
	CentraleModule []CentraleModule `gorm:"foreignKey:CarUserID"`
}

func NewCarUser(user_uuid uuid.UUID, name string) *Car_User {
	return &Car_User{
		User_uuid: user_uuid,
		Name: name,
	}
}

func (c *Car_User) GetID() uint {
	return c.ID
}

func (c *Car_User) GetUserUUID() uuid.UUID {
	return c.User_uuid
}

func (c *Car_User) GetName() string {
	return c.Name
}

func (c *Car_User) SetName(name string) {
	c.Name = name
}

func (c *Car_User) SetUserUUID(user_uuid uuid.UUID) {
	c.User_uuid = user_uuid
}

func (c *Car_User) SetCar(car uint) {
	c.CarID = car
}

func (c *Car_User) GetCar() uint {
	return c.CarID
}