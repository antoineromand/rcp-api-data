package contentcontroller

import (
	"encoding/json"
	"net/http"
	dto "rcp-api-data/internal/dto/content"
	usecase "rcp-api-data/internal/usecase/content"

	"gorm.io/gorm"
)

type MessageController struct {
	DB *gorm.DB
}

func NewMessageController(db *gorm.DB) *MessageController {
	return &MessageController{
		DB: db,
	}
}

func (c *MessageController) Controller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			// get message
			usecase := usecase.NewGetMessageUseCase(c.DB)
			result := usecase.GetLastMessage()
			if !result.Success {
				http.Error(w, result.Message, http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(result)
		case "POST":
			// get from json body the string content
			var content dto.MessageDTO
			err := json.NewDecoder(r.Body).Decode(&content)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			usecase := usecase.NewInsertMessageUseCase(c.DB)
			result := usecase.InsertMessage(content.Content)
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
