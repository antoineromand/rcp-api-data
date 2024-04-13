package carcontroller

import (
	"encoding/json"
	"net/http"
	dto "rcp-api-data/internal/dto/data-collector"
	usecase "rcp-api-data/internal/usecase/data-collector"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)

type CarUserController struct {
	DB *gorm.DB
}

func NewCarUserController(db *gorm.DB) *CarUserController {
	return &CarUserController{
		DB: db,
	}
}

func (cc *CarUserController) Controller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sugar := utils.GetLogger()
		token, err := utils.GetContextToken(r)
		if err != nil {
			http.Error(w, "Token introuvable", http.StatusUnauthorized)
		}
		switch r.Method {
		case "POST":
			var dto dto.CarUserDTO
			err := json.NewDecoder(r.Body).Decode(&dto)
			if err != nil {
				http.Error(w, "Error while decoding request body", http.StatusBadRequest)
				return
			}
			carUserBytes, err := json.Marshal(dto)
			if err != nil {
				http.Error(w, "Error while marshalling car user DTO", http.StatusBadRequest)
				return
			}
			usecase := usecase.NewInsertCarUserUseCase(cc.DB)

			response := usecase.InsertCarUser(token.UUID, carUserBytes)
			if !response.Success {
				sugar.Error(response.Message)
				http.Error(w, response.Message, response.Code)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(response)
		case "GET":
			usecase := usecase.NewGetCarsUserByIdUseCase(cc.DB)
			response := usecase.GetCarsUserById(token.UUID)
			if !response.Success {
				sugar.Error(response.Message)
				http.Error(w, response.Message, http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)

		}

	}
}
