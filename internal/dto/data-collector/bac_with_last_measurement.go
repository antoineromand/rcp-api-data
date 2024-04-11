package dto

type BacWithLastMeasurement struct {
	ID              uint    `json:"id"`
	Name            string  `json:"name"`
	LastMeasurement float64 `json:"last_measurement"`
}
