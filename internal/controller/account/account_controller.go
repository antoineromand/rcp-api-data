package controller_account

import (
	"encoding/json"
	"net/http"
	"rcp-api-data/internal/dto"
	"rcp-api-data/internal/usecase"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)

type AccountController struct {
	db *gorm.DB
}

func NewAccountController(db *gorm.DB) *AccountController {
	return &AccountController{
		db: db,
	}
}

func (c *AccountController) GetController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := utils.GetContextToken(r)
		if err != nil {
			http.Error(w, "Token introuvable", http.StatusUnauthorized)

		}
		var accountDTO dto.AccountDTO
		switch r.Method {
		case "GET":
			_usecase := usecase.NewGetInformationsUseCase(c.db)
			response := _usecase.GetInformationsByUserUuid(&token)
			if response.Error != nil {
				http.Error(w, response.Error.Message, response.Code)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response.Data)
		case "PUT":
			err := json.NewDecoder(r.Body).Decode(&accountDTO)
			if err != nil {
				http.Error(w, "Error while decoding request body", http.StatusBadRequest)
				return
			}
			accountBytes, err := json.Marshal(accountDTO)
			_usecase := usecase.NewPutInformationsUsecase(c.db)
			response := _usecase.PutInformations(token.UUID, accountBytes)
			if err != nil {
				http.Error(w, "Error while marshalling account DTO", http.StatusBadRequest)
				return
			}
			if response.Error != nil {
				http.Error(w, response.Error.Message, response.Code)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response.Data)
		case "DELETE":
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

	}
}
