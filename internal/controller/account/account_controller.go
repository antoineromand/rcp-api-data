package controller_account

import (
	"encoding/json"
	"net/http"
	"rcp-api-data/internal/dto"
	"rcp-api-data/internal/usecase"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)

func AccountController(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
		token, err := utils.GetContextToken(r)
		if err != nil {
			http.Error(w, "Token introuvable", http.StatusUnauthorized)
			
		}
		var accountDTO dto.AccountDTO
        switch r.Method {
			case "POST":
				err := json.NewDecoder(r.Body).Decode(&accountDTO)
				if err != nil {
					http.Error(w, "Error while decoding request body", http.StatusBadRequest)
				}
				response := usecase.CreateInformations(db, token.UUID, accountDTO)
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
				}
				response := usecase.CreateInformations(db, token.UUID, accountDTO)
				if response.Error != nil {
					http.Error(w, response.Error.Message, response.Code)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(response.Code)
				json.NewEncoder(w).Encode(response.Data)
			case "GET":
				response := usecase.GetInformationsByUserUuid(db, token.UUID)
				if response.Error != nil {
					http.Error(w, response.Error.Message, response.Code)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(response.Code)
				json.NewEncoder(w).Encode(response.Data)
		}
			
    }
}

