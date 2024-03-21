package utils

import "github.com/google/uuid"
func ConvertStringToUUID(s string) (uuid.UUID, error) {
    id, err := uuid.Parse(s)
    if err != nil {
        return uuid.Nil, err
    }
    return id, nil
}