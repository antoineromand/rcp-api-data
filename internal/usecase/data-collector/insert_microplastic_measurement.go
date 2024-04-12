package usecase

import (
	"fmt"
	"rcp-api-data/internal/entity/domain/data"
	mapper "rcp-api-data/internal/mapper/data-collector"
	repository "rcp-api-data/internal/repository/data-collector"
	"rcp-api-data/internal/utils"
	"strconv"

	"rcp-api-data/internal/config"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InsertMicroplasticMeasurementUseCase struct {
	DB *gorm.DB
}

type IMResponse struct {
	Success bool
	Code    int
	Message string
}

func NewInsertMicroplasticMeasurement(db *gorm.DB) *InsertMicroplasticMeasurementUseCase {
	return &InsertMicroplasticMeasurementUseCase{
		DB: db,
	}
}

func (e *InsertMicroplasticMeasurementUseCase) InsertMicroplasticMeasurement(_uuid string, bytes []byte) IMResponse {
	tx := e.DB.Begin()
	measures, err := mapper.MicroplasticMeasurementMapping(bytes)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error mapping microplastic measurement : ", err)
		return IMResponse{
			Success: false,
			Code:    500,
			Message: "Error mapping microplastic measurement",
		}
	}
	uuid, err := utils.ConvertStringToUUID(_uuid)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error converting string to uuid : ", err)
		return IMResponse{
			Success: false,
			Code:    500,
			Message: "Error converting string to uuid",
		}
	}
	// check if car user exist with user uuid and car_user_id :
	userExists := e.checkIfCarUserExists(uuid, measures.Car_User_ID)
	if !userExists {
		tx.Rollback()
		return IMResponse{
			Success: false,
			Code:    404,
			Message: "Car user not found",
		}
	}
	var moduleId *uint
	moduleId, err = e.checkIfModuleExists(measures.SSID)
	if err != nil {
		// Créer le module central
		moduleId, err = e.createModule(measures.SSID, "123456789", measures.Car_User_ID)
		if err != nil {
			tx.Rollback()
			fmt.Println("Error creating module")
			return IMResponse{
				Success: false,
				Code:    500,
				Message: "Erreur lors de la création du module",
			}
		}
	}

	// get bacs id in measures.Values
	for bacIDStr, value := range measures.Values {
		// Convertir bacIDStr en uint
		bacID, err := strconv.Atoi(bacIDStr)
		if err != nil {
			tx.Rollback()
			// Gérer l'erreur
			return IMResponse{
				Success: false,
				Code:    500,
				Message: "Erreur lors de la conversion de l'ID du BAC en entier",
			}
		}

		// Vérifier si le BAC existe
		bacExists := e.checkIfBacExists(uint(bacID))
		if !bacExists {
			created := e.createBac(uint(bacID), *moduleId)
			if !created {
				tx.Rollback()
				return IMResponse{
					Success: false,
					Code:    500,
					Message: "Erreur lors de la création du BAC",
				}
			}
		}

		// Créer la mesure de microplastique
		microplastic_measurement := data.MicroplasticMeasurement{
			BacID:  uint(bacID),
			Weight: value,
		}
		// Créer une instance de repository
		microplastic_measurement_repository := repository.NewMicroplasticMeasurementRepository(e.DB)

		// Créer la mesure de microplastique
		_, err = microplastic_measurement_repository.CreateMicroplasticMeasurement(&microplastic_measurement)
		if err != nil {
			tx.Rollback()

			fmt.Println("Erreur lors de la création de la mesure de microplastique : ", err)
			return IMResponse{
				Success: false,
				Code:    500,
				Message: "Erreur lors de la création de la mesure de microplastique",
			}
		}
	}

	tx.Commit()
	return IMResponse{
		Success: true,
		Code:    200,
		Message: "Microplastic measurement inserted",
	}
}

func (e *InsertMicroplasticMeasurementUseCase) checkIfCarUserExists(_uuid uuid.UUID, car_user_id uint) bool {
	car_user_repository := repository.NewCarUserRepository(e.DB)
	userExists := car_user_repository.CheckIfCarUserExistByUserUUIDAndCarUserID(_uuid, car_user_id)
	return userExists
}

func (e *InsertMicroplasticMeasurementUseCase) checkIfModuleExists(ssid string) (*uint, error) {
	module_repository := repository.NewModuleRepository(e.DB)
	module, err := module_repository.GetCentraleModuleByModuleSSID(ssid)
	if err != nil {
		return nil, err
	}
	return &module.ID, nil
}

func (e *InsertMicroplasticMeasurementUseCase) createModule(ssid, password string, car_user_id uint) (*uint, error) {
	hash, err := config.HashPassword(password)
	if err != nil {
		fmt.Println("Error hashing password : ", err)
	}
	module := data.NewCentraleModule(ssid, hash, car_user_id)
	module_repository := repository.NewModuleRepository(e.DB)
	id, err := module_repository.CreateCentraleModule(module)
	if err != nil {
		fmt.Println("Error creating module : ", err)
		return nil, err
	}
	return id, nil
}

func (e *InsertMicroplasticMeasurementUseCase) checkIfBacExists(bac_id uint) bool {
	bac_repository := repository.NewBacRepository(e.DB)
	_, err := bac_repository.GetBacByID(bac_id)
	return err == nil
}

func (e *InsertMicroplasticMeasurementUseCase) createBac(bac_id uint, module_id uint) bool {
	bac := data.NewBac(bac_id, module_id)
	bac_repository := repository.NewBacRepository(e.DB)
	_, err := bac_repository.CreateBac(bac)
	if err != nil {
		fmt.Println("Error creating bac : ", err)
		return false
	}
	return true
}
