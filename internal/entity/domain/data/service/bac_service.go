package service

import (
	entity "rcp-api-data/internal/entity/domain/data"
	"gorm.io/gorm"
)

type BacService struct {
	DB *gorm.DB
}


func NewBacService(db *gorm.DB) *BacService {
	return &BacService{
		DB: db,
	}
}

func (b *BacService) GetBacByID(id uint) (*entity.Bac, error) {
	var bac entity.Bac
	err := b.DB.Model(&entity.Bac{}).Where("id = ?", id).First(&bac).Error
	if err != nil {
		return nil, err
	}
	return &bac, nil
}




