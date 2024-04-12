package data

import "gorm.io/gorm"

type Bac struct {
	gorm.Model
	Id               uint                      `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name             *string                   `json:"name"`
	CentraleModuleID uint                      `json:"centrale_module_id"`
	Mesurements      []MicroplasticMeasurement `gorm:"foreignKey:BacID"`
}

func (Bac) TableName() string {
	return "bac"
}

func NewBac(id uint, central_module uint) *Bac {
	return &Bac{
		Id:               id,
		CentraleModuleID: central_module,
	}
}
