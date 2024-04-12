package repository

import (
	"rcp-api-data/internal/entity/domain/data"

	"gorm.io/gorm"
)

type ICarRepository interface {
	CreateCar(*data.Car) (bool, error)
	GetCarByID(uint) (*data.Car, error)
	GetAllCars() ([]data.Car, error)
	UpdateCar(*data.Car) (*data.Car, error)
	DeleteCarByID(uint) error
}

type CarRepository struct {
	DB *gorm.DB
	ICarRepository
}

func NewCarRepository(db *gorm.DB) ICarRepository {
	return &CarRepository{
		DB: db,
	}
}

func (cr *CarRepository) CreateCar(car *data.Car) (bool, error) {
	if err := cr.DB.Create(car).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (cr *CarRepository) GetCarByID(id uint) (*data.Car, error) {
	car := &data.Car{}
	result := cr.DB.Where("id = ?", id).First(car)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return car, nil
}

func (cr *CarRepository) UpdateCar(car *data.Car) (*data.Car, error) {
	if err := cr.DB.Save(car).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func (cr *CarRepository) GetAllCars() ([]data.Car, error) {
	var cars []data.Car
	result := cr.DB.Find(&cars)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return cars, nil
}

func (cr *CarRepository) DeleteCarByID(id uint) error {
	if err := cr.DB.Where("id = ?", id).Delete(&data.Car{}).Error; err != nil {
		return err
	}
	return nil
}
