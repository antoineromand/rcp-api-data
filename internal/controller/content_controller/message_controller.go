package contentcontroller

import (
	"encoding/json"
	"net/http"
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
			// insert message
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
