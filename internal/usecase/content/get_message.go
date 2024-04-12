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

func (e *GetMessageUseCase) GetMessage(id uint) (entity_content.Message, error) {
	res, err := repository.NewMessageRepository(e.DB).GetMessage(id)
	if err != nil {
		return entity_content.Message{}, err
	}
	return res, nil
}
