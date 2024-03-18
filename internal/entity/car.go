package entity

import (
	"time"
)

type FuelType string

const (
	ELECTRICAL FuelType = "ELECTRICAL"
	GASOLINE   FuelType = "GASOLINE"
	DIESEL     FuelType = "DIESEL"
)

type Car struct {
	ID uint `json:"id" gorm:"primary_key AUTO_INCREMENT"`
	CarBrandID uint `json:"brand_id"`
	Year int `json:"year"`
	FuelType FuelType `json:"fuel_type" gorm:"type:fuelType"`
	Car_Model string `json:"model"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func NewCar(brand uint, year int, fuelType FuelType, model string) *Car {
	return &Car{
		CarBrandID: brand,
		Year: year,
		FuelType: fuelType,
		Car_Model: model,
	}
}

func (c *Car) GetID() uint {
	return c.ID
}

func (c *Car) GetBrand() uint {
	return c.CarBrandID
}

func (c *Car) GetYear() int {
	return c.Year
}

func (c *Car) GetFuelType() FuelType {
	return c.FuelType
}

func (c *Car) GetModel() string {
	return c.Car_Model
}

func (c *Car) SetBrand(brand uint) {
	c.CarBrandID = brand
}

func (c *Car) SetYear(year int) {
	c.Year = year
}

func (c *Car) SetFuelType(fuelType FuelType) {
	c.FuelType = fuelType
}

func (c *Car) SetModel(model string) {
	c.Car_Model = model
}

