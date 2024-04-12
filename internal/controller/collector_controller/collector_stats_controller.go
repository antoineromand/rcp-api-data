package collectorcontroller

import (
	"encoding/json"
	"net/http"
	usecase "rcp-api-data/internal/usecase/data-collector"
	"rcp-api-data/internal/utils"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type CollectorStatsController struct {
	DB *gorm.DB
}

func NewCollectorStatsController(db *gorm.DB) *CollectorStatsController {
	return &CollectorStatsController{
		DB: db,
	}
}

func (cc *CollectorStatsController) Controller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token, err := utils.GetContextToken(r)
		if err != nil {
			http.Error(w, "Token introuvable", http.StatusUnauthorized)
		}
		switch r.Method {
		case "GET":
			param := strings.Replace(r.URL.Path, "/api/recup-plast/data-collector/stats/", "", 1)
			// check if param is uint
			// if not, return error
			param = strings.TrimSpace(param)
			if param == "" {
				http.Error(w, "Invalid parameter", http.StatusBadRequest)
				return
			}
			paramUint, err := strconv.ParseUint(param, 10, 64)
			if err != nil {
				http.Error(w, "Invalid parameter", http.StatusBadRequest)
				return
			}
			usecase := usecase.NewGetDataCollectorUseCase(cc.DB)
			result := usecase.GetDataByBacs(token.UUID, uint(paramUint))
			if !result.Success {
				http.Error(w, result.Message, http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(result)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
