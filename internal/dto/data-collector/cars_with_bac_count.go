package dto

type CarWithBacCount struct {
	ID       uint   `json:"id"`
	BrandName  uint   `json:"brand_name"`
	Year     int    `json:"year"`
	FuelType string `json:"fuel_type"`
	Model    string `json:"model"`
	BacCount uint64 `json:"bac_count"`
}
