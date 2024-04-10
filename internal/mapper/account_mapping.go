package mapper

import (
	"encoding/json"
	"errors"
	"rcp-api-data/internal/dto"
	"rcp-api-data/internal/entity/domain/account"

	"github.com/google/uuid"
)

type AccountMappingResponse struct {
	Account  account.Account
	Password *string
}

func AccountMapping(dtoBytes []byte, uuid *uuid.UUID) (AccountMappingResponse, error) {
	var dto dto.AccountDTO
	if err := json.Unmarshal(dtoBytes, &dto); err != nil {
		return AccountMappingResponse{}, err
	}

	if uuid == nil {
		return AccountMappingResponse{}, errors.New("UUID is nil")
	}
	return AccountMappingResponse{
		Account: account.Account{
			ActivityMessage: dto.ActivityMessage,
			Address:         dto.Address,
			City:            dto.City,
			Country:         dto.Country,
			PostalCode:      dto.PostalCode,
			PhoneNumber:     dto.PhoneNumber,
			FirstName:       dto.FirstName,
			LastName:        dto.LastName,
			IsNew:           dto.IsNew,
			Username:        dto.Username,
			Email:           dto.Email,
			UserUUID:        *uuid,
		},
		Password: dto.Pasword,
	}, nil
}
