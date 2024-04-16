package data

import "gorm.io/gorm"

type Bac struct {
	gorm.Model
	EID              uint                      `json:"eid"`
	CentraleModuleID uint                      `json:"centrale_module_id"`
	Mesurements      []MicroplasticMeasurement `gorm:"foreignKey:BacID"`
}

func (Bac) TableName() string {
	return "bac"
}

func NewBac(eid uint, central_module uint) *Bac {
	return &Bac{
		EID:              eid,
		CentraleModuleID: central_module,
	}
}
