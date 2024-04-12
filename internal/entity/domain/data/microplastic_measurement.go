package data

import (
	"gorm.io/gorm"
)

type MicroplasticMeasurement struct {
	gorm.Model
	ID     uint `gorm:"primaryKey;autoIncrement"`
	BacID  uint
	Weight uint
}

func NewMicroplasticMeasurement(bac uint, weight uint) *MicroplasticMeasurement {
	return &MicroplasticMeasurement{
		BacID:  bac,
		Weight: weight,
	}
}

func (MicroplasticMeasurement) TableName() string {
	return "microplastic_measurement"
}
