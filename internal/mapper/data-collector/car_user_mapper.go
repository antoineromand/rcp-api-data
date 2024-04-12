package mapper

import (
	"encoding/json"
	dto "rcp-api-data/internal/dto/data-collector"
	"rcp-api-data/internal/utils"
)

func CarUserMapping(dtoBytes []byte, uuid string) (dto.CarUserDTO, error) {
	var car_user dto.CarUserDTO
	if err := json.Unmarshal(dtoBytes, &car_user); err != nil {
		return dto.CarUserDTO{}, err
	}
	_uuid, err := utils.ConvertStringToUUID(uuid)
	if err != nil {
		return dto.CarUserDTO{}, err
	}
	car_user.UserUUID = _uuid
	return car_user, nil
}
