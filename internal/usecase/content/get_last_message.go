package usecase

import (
	entity_content "rcp-api-data/internal/entity/domain/content"
	repository "rcp-api-data/internal/repository/content"

	"gorm.io/gorm"
)

type GetMessageUseCase struct {
	DB *gorm.DB
}

func NewGetMessageUseCase(db *gorm.DB) *GetMessageUseCase {
	return &GetMessageUseCase{
		DB: db,
	}
}

type IGMResponse struct {
	Success bool
	Message string
	Data    entity_content.Message
}

func (e *GetMessageUseCase) GetLastMessage() IGMResponse {
	res, err := repository.NewMessageRepository(e.DB).GetMessage()

	if err != nil {
		return IGMResponse{
			Success: false,
			Message: err.Error(),
			Data:    entity_content.Message{},
		}
	}

	return IGMResponse{
		Success: true,
		Message: "Message fetched successfully",
		Data:    res,
	}
}
