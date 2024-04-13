package usecase

import (
	repository "rcp-api-data/internal/repository/content"

	"gorm.io/gorm"
)

type InsertMessageUseCase struct {
	DB *gorm.DB
}

type IIMResponse struct {
	Success bool
	Message string
}

func NewInsertMessageUseCase(db *gorm.DB) *InsertMessageUseCase {
	return &InsertMessageUseCase{
		DB: db,
	}
}

func (e *InsertMessageUseCase) InsertMessage(content string) IIMResponse {
	err := repository.NewMessageRepository(e.DB).InsertMessage(content)

	if err != nil {
		return IIMResponse{
			Success: false,
			Message: err.Error(),
		}
	}
	return IIMResponse{
		Success: true,
		Message: "Message inserted successfully",
	}
}
