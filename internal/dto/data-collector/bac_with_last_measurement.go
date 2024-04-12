package dto

type BacWithLastMeasurement struct {
	ID              uint    `json:"id"`
	Date            string  `json:"date"`
	LastMeasurement float64 `json:"last_measurement"`
}
