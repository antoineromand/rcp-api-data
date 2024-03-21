package entity

import (
	"time"

	"github.com/google/uuid"
)

type CentraleModule struct {
	ID        uuid.UUID 
	Name      string     
	SSID      string    
	Password  string     
	Safer     string     
	Car_User  Car_User   `gorm:"foreignKey:CarUserID"`
	CarUserID uint       
	CreatedAt time.Time  
	Bac []Bac `gorm:"foreignKey:CentraleModuleID"`
}

func NewCentraleModul(name, ssid, password, safer string, car_user Car_User) *CentraleModule {
	return &CentraleModule{
		ID:        uuid.New(),
		Name:      name,
		SSID:      ssid,
		Password:  password,
		Safer:     safer,
		Car_User:  car_user,
		CreatedAt: time.Now(),
	}
}

func (c *CentraleModule) GetID() uuid.UUID {
	return c.ID
}

func (c *CentraleModule) GetName() string {
	return c.Name
}

func (c *CentraleModule) GetSSID() string {
	return c.SSID
}

func (c *CentraleModule) GetPassword() string {
	return c.Password
}

func (c *CentraleModule) GetSafer() string {
	return c.Safer
}

func (c *CentraleModule) GetCarUser() Car_User {
	return c.Car_User
}

func (c *CentraleModule) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c *CentraleModule) SetName(name string) {
	c.Name = name
}

func (c *CentraleModule) SetSSID(ssid string) {
	c.SSID = ssid
}

func (c *CentraleModule) SetPassword(password string) {
	c.Password = password
}

func (c *CentraleModule) SetSafer(safer string) {
	c.Safer = safer
}

func (c *CentraleModule) SetCarUser(car_user Car_User) {
	c.Car_User = car_user
}

func (c *CentraleModule) GetDateToString() string {
	return c.CreatedAt.String()
}
