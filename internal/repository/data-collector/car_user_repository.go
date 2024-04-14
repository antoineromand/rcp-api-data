package repository

import (
	dto "rcp-api-data/internal/dto/data-collector"
	"rcp-api-data/internal/entity/domain/data"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ICarUserRepository interface {
	CreateCarUser(*data.Car_User) (*data.Car_User, error)
	GetCarUserByID(uint) (*data.Car_User, error)
	GetAllCarUsers() ([]data.Car_User, error)
	GetAllCarByUserUuid(uuid.UUID) ([]data.Car_User, error)
	UpdateCarUser(*data.Car_User) (*data.Car_User, error)
	DeleteCarUserByID(uint) error
	CheckIfCarUserExistByUserUUIDAndCarUserID(uuid.UUID, uint) bool
	GetCarsWithBacCountByUserUUID(uuid.UUID) ([]dto.CarWithBacCount, error)
}

type CarUserRepository struct {
	DB *gorm.DB
	ICarUserRepository
}

func NewCarUserRepository(db *gorm.DB) ICarUserRepository {
	return &CarUserRepository{
		DB: db,
	}
}

func (cur *CarUserRepository) CreateCarUser(car_user *data.Car_User) (*data.Car_User, error) {
	if err := cur.DB.Create(car_user).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func (cur *CarUserRepository) GetCarUserByID(id uint) (*data.Car_User, error) {
	car_user := &data.Car_User{}
	result := cur.DB.Where("id = ?", id).First(car_user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return car_user, nil
}

func (cur *CarUserRepository) GetCarUserByUserID(id uuid.UUID) (*data.Car_User, error) {
	car_user := &data.Car_User{}
	result := cur.DB.Where("id = ?", id).First(car_user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return car_user, nil
}

func (cur *CarUserRepository) UpdateCarUser(car_user *data.Car_User) (*data.Car_User, error) {
	if err := cur.DB.Save(car_user).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func (cur *CarUserRepository) GetAllCarUsers() ([]data.Car_User, error) {
	var car_users []data.Car_User
	result := cur.DB.Find(&car_users)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return car_users, nil
}

func (cur *CarUserRepository) GetAllCarByUserUuid(uuid uuid.UUID) ([]data.Car_User, error) {
	var car_users []data.Car_User
	result := cur.DB.Where("user_uuid = ?", uuid).Find(&car_users)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return car_users, nil
}

func (cur *CarUserRepository) DeleteCarUserByID(id uint) error {
	result := cur.DB.Where("id = ?", id).Delete(&data.Car_User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cur *CarUserRepository) CheckIfCarUserExistByUserUUIDAndCarUserID(userUUID uuid.UUID, carUserID uint) bool {
	var count int64
	cur.DB.Model(&data.Car_User{}).Where("user_uuid = ? AND id = ?", userUUID, carUserID).Count(&count)
	return count > 0
}

func (cur *CarUserRepository) GetCarsWithBacCountByUserUUID(userUUID uuid.UUID) ([]dto.CarWithBacCount, error) {
	var cars []dto.CarWithBacCount
	result := cur.DB.
		Model(&data.Car{}).
		Joins("JOIN car_user ON car.id = car_user.car_id").
		Joins("JOIN brand ON car.car_brand_id = brand.id").
		Joins("LEFT JOIN centrale_module ON car_user.id = centrale_module.car_user_id").
		Joins("LEFT JOIN bac ON centrale_module.id = bac.centrale_module_id").
		Where("car_user.user_uuid = ?", userUUID).
		Select("car_user.id AS id, brand.name AS brand, car.year, car.fuel_type, car.car_model AS model, COUNT(bac.id) AS bac_count").
		Group("car_user.id, brand.name, car.year, car.fuel_type, car.car_model").
		Find(&cars)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return cars, nil
}
