package repository

import (
	"rcp-api-data/internal/entity/domain/data"

	"gorm.io/gorm"
)

type IBrandCrudRepository interface {
	CreateBrand(*data.Brand) (*data.Brand, error)
	GetBrandByID(uint) (*data.Brand, error)
	GetAllBrands() ([]data.Brand, error)
	UpdateBrand(*data.Brand) (*data.Brand, error)
	DeleteBrandByID(uint) error
}

type BrandCrudRepository struct {
	DB *gorm.DB
	IBrandCrudRepository
}

func NewBrandCrudRepository(db *gorm.DB) IBrandCrudRepository {
	return &BrandCrudRepository{
		DB: db,
	}
}

func (bcr *BrandCrudRepository) CreateBrand(brand *data.Brand) (*data.Brand, error) {
	if err := bcr.DB.Create(brand).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func (bcr *BrandCrudRepository) GetBrandByID(id uint) (*data.Brand, error) {
	brand := &data.Brand{}
	result := bcr.DB.Where("id = ?", id).First(brand)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return brand, nil
}

func (bcr *BrandCrudRepository) UpdateBrand(brand *data.Brand) (*data.Brand, error) {
	if err := bcr.DB.Save(brand).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func (bcr *BrandCrudRepository) GetAllBrands() ([]data.Brand, error) {
	var brands []data.Brand
	result := bcr.DB.Find(&brands)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return brands, nil
}

func (bcr *BrandCrudRepository) DeleteBrandByID(id uint) error {
	if err := bcr.DB.Delete(&data.Brand{}, id).Error; err != nil {
		return err
	}
	return nil
}