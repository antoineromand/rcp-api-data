package entity

import (
	"time"
)

type MicroplasticMeasurement struct {
    ID uint `gorm:"primaryKey;autoIncrement"`
	BacID uint 
	Weight float64
	CreatedAt time.Time 
}

func NewMicroplasticMeasurement(bac uint, weight float64) *MicroplasticMeasurement {
	return &MicroplasticMeasurement{
		BacID: bac,
		Weight: weight,
		CreatedAt: time.Now(),
	}
}

func (m *MicroplasticMeasurement) GetID() uint {
	return m.ID
}

func (m *MicroplasticMeasurement) GetBacID() uint {
	return m.BacID
}

func (m *MicroplasticMeasurement) GetWeight() float64 {
	return m.Weight
}

func (m *MicroplasticMeasurement) GetCreatedAt() time.Time {
	return m.CreatedAt
}

func (m *MicroplasticMeasurement) SetBac(bac uint) {
	m.BacID = bac
}

func (m *MicroplasticMeasurement) SetWeight(weight float64) {
	m.Weight = weight
}

func (m *MicroplasticMeasurement) SetCreatedAt(created_at time.Time) {
	m.CreatedAt = created_at
}