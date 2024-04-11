package data

import "time"

type Brand struct {
	ID        uint      `json:"id" gorm:"primary_key AUTO_INCREMENT"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Cars      []Car     `gorm:"foreignKey:CarBrandID"`
}

func NewBrand(name string) *Brand {
	return &Brand{
		Name: name,
	}
}

func (Brand) TableName() string {
	return "brand"
}

func (b *Brand) GetID() uint {
	return b.ID
}

func (b *Brand) GetName() string {
	return b.Name
}
