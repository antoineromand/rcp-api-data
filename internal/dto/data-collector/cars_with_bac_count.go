package dto

type CarWithBacCount struct {
	ID       uint   `json:"id"`
	Brand    string `json:"brand"`
	Year     int    `json:"year"`
	FuelType string `json:"fuel_type"`
	Model    string `json:"model"`
	BacCount uint64 `json:"nb_bacs"`
}
