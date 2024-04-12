package data

import (
	"gorm.io/gorm"
)

type CentraleModule struct {
	gorm.Model
	Name      *string  `gorm:"null"`
	SSID      string   `json:"ssid" gorm:"name:ssid;unique;not null"`
	Password  string   `json:"-"`
	Safer     *string  `gorm:"null"`
	CarUserID uint     `gorm:"unique"`
	Car_User  Car_User `gorm:"foreignKey:CarUserID"`
	Bac       []Bac    `gorm:"foreignKey:CentraleModuleID"`
}

func (CentraleModule) TableName() string {
	return "centrale_module"
}

func NewCentraleModule(ssid, password string, car_user uint) *CentraleModule {
	return &CentraleModule{
		SSID:      ssid,
		Password:  password,
		CarUserID: car_user,
	}
}
