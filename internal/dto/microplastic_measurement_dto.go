package dto

import "errors"

type MicroplasticMeasurementDTO struct {
	BacID     uint   `json:"bac_id"`
	Weight    float64 `json:"weight"`
	CreatedAt string `json:"created_at"`
}

func NewMicroplasticMeasurementDTO(bac *uint, weight *float64, created_at *string) (*MicroplasticMeasurementDTO, error) {
	if bac == nil {
		return nil, errors.New("please provide a valid bac")
	}
	if weight == nil {
		return nil, errors.New("please provide a valid weight")
	}
	if created_at == nil {
		return nil, errors.New("please provide a valid created_at")
	}
	return &MicroplasticMeasurementDTO{
		BacID:     *bac,
		Weight:    *weight,
		CreatedAt: *created_at,
	}, nil
}
