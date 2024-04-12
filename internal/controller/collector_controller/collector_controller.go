package collectorcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	dto "rcp-api-data/internal/dto/data-collector"
	usecase "rcp-api-data/internal/usecase/data-collector"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)

type CollectorController struct {
	DB *gorm.DB
}

func NewCollectorController(db *gorm.DB) *CollectorController {
	return &CollectorController{
		DB: db,
	}
}

func (cc *CollectorController) Controller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := utils.GetContextToken(r)
		if err != nil {
			http.Error(w, "Token introuvable", http.StatusUnauthorized)
		}
		switch r.Method {
		case "POST":
			var dto dto.MicroplasticMeasurement
			err := json.NewDecoder(r.Body).Decode(&dto)
			if err != nil {
				fmt.Println("Error while decoding request body : ", err)
				http.Error(w, "Error while decoding request body", http.StatusBadRequest)
				return
			}
			measure, err := json.Marshal(dto)
			if err != nil {
				http.Error(w, "Error while marshalling car user DTO", http.StatusBadRequest)
				return
			}
			usecase := usecase.NewInsertMicroplasticMeasurement(cc.DB)
			response := usecase.InsertMicroplasticMeasurement(token.UUID, measure)
			if !response.Success {
				http.Error(w, response.Message, response.Code)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
