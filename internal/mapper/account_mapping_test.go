package mapper

import (
	"encoding/json"
	"rcp-api-data/internal/dto"
	"testing"

	"github.com/google/uuid"
)

func TestAccountMapping(t *testing.T) {
	t.Run("should not return error", func(t *testing.T) {
		// Créer un objet JSON d'input avec des valeurs null pour certains champs
		inputData := dto.AccountDTO{
			ActivityMessage: nil,
			Address:         nil,
			City:            nil,
			Country:         nil,
			PostalCode:      nil,
			PhoneNumber:     nil,
			FirstName:       nil,
			LastName:        nil,
			IsNew:           nil,
		}
		uuid := uuid.New()

		// Convertir inputData en un tableau d'octets (byte array)
		inputDataBytes, err := json.Marshal(inputData)
		if err != nil {
			t.Errorf("Error while marshalling input JSON: %v", err)
		}

		account, err := AccountMapping(inputDataBytes, &uuid)
		if err != nil {
			t.Errorf("Error while creating input JSON: %v", err)
		}

		if account.UserUUID != uuid {
			t.Errorf("Error while mapping account")
		}
	})

	t.Run("should return error if json is invalid", func(t *testing.T) {
		invalidJSON := []byte(`{invalid json}`)
		uuid := uuid.New()

		_, err := AccountMapping(invalidJSON, &uuid)
		if err == nil {
			t.Errorf("Expected unmarshal error, got nil")
		}
	})

	// Test for nil UUID
	t.Run("should return an error if uuid is nil", func(t *testing.T) {
		inputData := dto.AccountDTO{
			ActivityMessage: nil,
			Address:         nil,
			City:            nil,
			Country:         nil,
			PostalCode:      nil,
			PhoneNumber:     nil,
			FirstName:       nil,
			LastName:        nil,
			IsNew:           nil,
		}
		inputDataBytes, err := json.Marshal(inputData)
		if err != nil {
			t.Errorf("Error marshalling input data: %v", err)
		}

		_, err = AccountMapping(inputDataBytes, nil)
		if err == nil {
			t.Errorf("UUID is nil")
		}
	})
}
