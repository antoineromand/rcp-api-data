package migration

import (
	"errors"
	"fmt"
	entity_account "rcp-api-data/internal/entity/domain/account"
	entity_data "rcp-api-data/internal/entity/domain/data"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)

// ErrMigrationNilDB is returned when the database connection is missing
var ErrMigrationNilDB = errors.New("cannot run migration, db connection is missing")

// RunMigration runs the database migration
func RunMigration(db *gorm.DB) error {
	sugar := utils.GetLogger()

	if db == nil {
		sugar.Panic(ErrMigrationNilDB)
		return ErrMigrationNilDB
	}

	if err := drop(db); err != nil {
		sugar.Errorw("failed to drop tables", "error", err)
		return fmt.Errorf("failed to drop tables: %w", err)
	}
	sugar.Info("Tables dropped")
	sugar.Info("Creating tables...")

	if err := createEnums(db); err != nil {
		sugar.Errorw("failed to create enums", "error", err)
		return fmt.Errorf("failed to create enums: %w", err)
	}
	sugar.Info("Enums created")

	if err := createTables(db); err != nil {
		sugar.Errorw("failed to create tables", "error", err)
		return fmt.Errorf("failed to create tables: %w", err)
	}
	sugar.Info("Tables created")

	return nil
}

// createEnums creates the fueltype enum in the database
func createEnums(db *gorm.DB) error {
	if db == nil {
		return nil
	}
	result := db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'fueltype') THEN
				CREATE TYPE fueltype AS ENUM ('electrical', 'gasoline', 'diesel');
			END IF;
		END $$;
	`)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// drop drops all tables from the database
func drop(db *gorm.DB) error {
	sugar := utils.GetLogger()
	if db == nil {
		return nil
	}
	if db.Migrator().HasTable(&entity_data.Brand{}) {
		if err := db.Migrator().DropTable(&entity_data.Brand{}); err != nil {
			sugar.Errorw("failed to drop brand table", "error", err)
			return fmt.Errorf("failed to drop brand table: %w", err)
		}
	}
	if db.Migrator().HasTable(&entity_data.Car{}) {
		if err := db.Migrator().DropTable(&entity_data.Car{}); err != nil {
			sugar.Errorw("failed to drop car table", "error", err)
			return fmt.Errorf("failed to drop car table: %w", err)
		}
	}
	if db.Migrator().HasTable(&entity_data.Car_User{}) {
		if err := db.Migrator().DropTable(&entity_data.Car_User{}); err != nil {
			sugar.Errorw("failed to drop car_user table", "error", err)
			return fmt.Errorf("failed to drop car_user table: %w", err)
		}
	}
	if db.Migrator().HasTable(&entity_data.CentraleModule{}) {
		if err := db.Migrator().DropTable(&entity_data.CentraleModule{}); err != nil {
			sugar.Errorw("failed to drop centrale_module table", "error", err)
			return fmt.Errorf("failed to drop centrale_module table: %w", err)
		}
	}
	if db.Migrator().HasTable(&entity_data.Bac{}) {
		if err := db.Migrator().DropTable(&entity_data.Bac{}); err != nil {
			sugar.Errorw("failed to drop bac table", "error", err)
			return fmt.Errorf("failed to drop bac table: %w", err)
		}
	}
	if db.Migrator().HasTable(&entity_data.MicroplasticMeasurement{}) {
		if err := db.Migrator().DropTable(&entity_data.MicroplasticMeasurement{}); err != nil {
			sugar.Errorw("failed to drop microplastic_measurement table", "error", err)
			return fmt.Errorf("failed to drop microplastic_measurement table: %w", err)
		}
	}
	return nil
}

// createTables creates all tables in the database
func createTables(db *gorm.DB) error {
	sugar := utils.GetLogger()
	if db == nil {
		return nil
	}
	if err := db.Migrator().CreateTable(&entity_data.Brand{}); err != nil {
		sugar.Errorw("failed to create brand table", "error", err)
		return fmt.Errorf("failed to create brand table: %w", err)
	}
	if err := db.Migrator().CreateTable(&entity_data.Car{}); err != nil {
		sugar.Errorw("failed to create car table", "error", err)
		return fmt.Errorf("failed to create car table: %w", err)
	}
	if err := db.Migrator().CreateTable(&entity_data.Car_User{}); err != nil {
		sugar.Errorw("failed to create car_user table", "error", err)
		return fmt.Errorf("failed to create car_user table: %w", err)
	}
	if err := db.Migrator().CreateTable(&entity_data.CentraleModule{}); err != nil {
		sugar.Errorw("failed to create centrale_module table", "error", err)
		return fmt.Errorf("failed to create centrale_module table: %w", err)
	}
	if err := db.Migrator().CreateTable(&entity_data.Bac{}); err != nil {
		sugar.Errorw("failed to create bac table", "error", err)
		return fmt.Errorf("failed to create bac table: %w", err)
	}
	if err := db.Migrator().CreateTable(&entity_data.MicroplasticMeasurement{}); err != nil {
		sugar.Errorw("failed to create microplastic_measurement table", "error", err)
		return fmt.Errorf("failed to create microplastic_measurement table: %w", err)
	}
	if err := db.Migrator().CreateTable(&entity_account.Account{}); err != nil {
		sugar.Errorw("failed to create user table", "error", err)
		return fmt.Errorf("failed to create user table: %w", err)
	}
	return nil
}
