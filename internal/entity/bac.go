package entity

import "gorm.io/gorm"

type Bac struct {
	gorm.Model
	Id uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name string 
	CentraleModuleID uint 
	Mesurements []MicroplasticMeasurement `gorm:"foreignKey:BacID"`
}

func NewBac(name string, central_module uint) *Bac {
	return &Bac{
		Name: name,
		CentraleModuleID: central_module,
	}
}

func (b *Bac) GetName() string {
	return b.Name
}

func (b *Bac) GetCentralModuleID() uint {
	return b.CentraleModuleID
}

func (b *Bac) SetName(name string) {
	b.Name = name
}

func (b *Bac) SetCentralModuleID(central_module uint) {
	b.CentraleModuleID = central_module
}
