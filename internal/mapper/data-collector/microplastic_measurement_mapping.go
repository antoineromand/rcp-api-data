package mapper

import (
	"encoding/json"
	dto "rcp-api-data/internal/dto/data-collector"
)

func MicroplasticMeasurementMapping(dtoBytes []byte) (dto.MicroplasticMeasurement, error) {
	var measures dto.MicroplasticMeasurement
	if err := json.Unmarshal(dtoBytes, &measures); err != nil {
		return dto.MicroplasticMeasurement{}, nil
	}

	return measures, nil
}
