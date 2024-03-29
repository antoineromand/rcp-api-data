package mapper

import (
	"encoding/json"
	"errors"
	"rcp-api-data/internal/dto"
	"rcp-api-data/internal/entity/domain/account"

	"github.com/google/uuid"
)

func AccountMapping(dtoBytes []byte, uuid *uuid.UUID) (account.Account, error) {
	var dto dto.AccountDTO
	if err := json.Unmarshal(dtoBytes, &dto); err != nil {
		return account.Account{}, err
	}

	if uuid == nil {
		return account.Account{}, errors.New("UUID is nil")
	}
	return account.Account{
		ActivityMessage: dto.ActivityMessage,
		Address:         dto.Address,
		City:            dto.City,
		Country:         dto.Country,
		PostalCode:      dto.PostalCode,
		PhoneNumber:     dto.PhoneNumber,
		FirstName:       dto.FirstName,
		LastName:        dto.LastName,
		IsNew:           dto.IsNew,
		UserUUID:        *uuid,
	}, nil
}
