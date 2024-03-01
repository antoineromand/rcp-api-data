package config

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

type TokenFromCookie struct {
    UUID string `json:"uuid"`
    Username string `json:"username"`
    RolePermission map[string][] string `json:"rolePermission"`
}



func DecodePayload(token string) (TokenFromCookie, error) {
	var tokenData TokenFromCookie
	splitToken := strings.Split(token, ".")
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
