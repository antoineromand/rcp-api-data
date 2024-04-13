package dto

type CarDTO struct {
	ID       uint   `json:"id"`
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	Year     int    `json:"year"`
	FuelType string `json:"fuel_type"`
}
