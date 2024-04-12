package carcontroller

import (
	"encoding/json"
	"net/http"
	usecase "rcp-api-data/internal/usecase/data-collector"
	"rcp-api-data/internal/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CarController struct {
	DB    *gorm.DB
	sugar *zap.SugaredLogger
}

func NewCarController(db *gorm.DB) *CarController {
	return &CarController{
		DB: db,
	}
}

func (cc *CarController) CarController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sugar := utils.GetLogger()

		switch r.Method {
		case "GET":
			usecase := usecase.NewGetCarsUseCase(cc.DB)
			response := usecase.GetCars()
			if !response.Success {
				sugar.Error("Error while getting cars")
				http.Error(w, "Error while getting cars", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response.Data)
		}
	}
}
