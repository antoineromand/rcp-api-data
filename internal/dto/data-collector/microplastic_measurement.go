package dto

type MicroplasticMeasurement struct {
	SSID        string          `json:"ssid"`
	Car_User_ID uint            `json:"car_user_id"`
	Values      map[string]uint `json:"values"`
}
