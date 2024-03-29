package security

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

type TokenFromCookie struct {
	UUID           string              `json:"uuid"`
	Username       string              `json:"username"`
	Email          string              `json:"email"`
	RolePermission map[string][]string `json:"rolePermission"`
}

func DecodePayload(token string) (TokenFromCookie, error) {
	if token == "" {
		return TokenFromCookie{}, errors.New("token is empty")
	}
	var tokenData TokenFromCookie
	splitToken := strings.Split(token, ".")
	if len(splitToken) < 2 {
		return tokenData, errors.New("invalid token")
	}
	payload, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		return tokenData, err
	}
	err = json.Unmarshal(payload, &tokenData)
	if err != nil {
		return tokenData, err
	}
	return tokenData, nil
}
