package migration

import (
	"errors"
	"fmt"
	"rcp-api-data/internal/entity"
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
	if db == nil {
		return nil
	}
	if db.Migrator().HasTable(&entity.Brand{}) {
		if err := db.Migrator().DropTable(&entity.Brand{}); err != nil {
			return fmt.Errorf("failed to drop brand table: %w", err)
		}
	}
	if db.Migrator().HasTable(&entity.Car{}) {
		if err := db.Migrator().DropTable(&entity.Car{}); err != nil {
			return fmt.Errorf("failed to drop car table: %w", err)
		}
	}
	if db.Migrator().HasTable(&entity.Car_User{}) {
		if err := db.Migrator().DropTable(&entity.Car_User{}); err != nil {
			return fmt.Errorf("failed to drop car_user table: %w", err)
		}
	}
	if db.Migrator().HasTable(&entity.CentraleModule{}) {
		if err := db.Migrator().DropTable(&entity.CentraleModule{}); err != nil {
			return fmt.Errorf("failed to drop centrale_module table: %w", err)
		}
	}
	if db.Migrator().HasTable(&entity.Bac{}) {
		if err := db.Migrator().DropTable(&entity.Bac{}); err != nil {
			return fmt.Errorf("failed to drop bac table: %w", err)
		}
	}
	if db.Migrator().HasTable(&entity.MicroplasticMeasurement{}) {
		if err := db.Migrator().DropTable(&entity.MicroplasticMeasurement{}); err != nil {
			return fmt.Errorf("failed to drop microplastic_measurement table: %w", err)
		}
	}
	return nil
}

// createTables creates all tables in the database
func createTables(db *gorm.DB) error {
	if db == nil {
		return nil
	}
	if err := db.Migrator().CreateTable(&entity.Brand{}); err != nil {
		return fmt.Errorf("failed to create brand table: %w", err)
	}
	if err := db.Migrator().CreateTable(&entity.Car{}); err != nil {
		return fmt.Errorf("failed to create car table: %w", err)
	}
	if err := db.Migrator().CreateTable(&entity.Car_User{}); err != nil {
		return fmt.Errorf("failed to create car_user table: %w", err)
	}
	if err := db.Migrator().CreateTable(&entity.CentraleModule{}); err != nil {
		return fmt.Errorf("failed to create centrale_module table: %w", err)
	}
	if err := db.Migrator().CreateTable(&entity.Bac{}); err != nil {
		return fmt.Errorf("failed to create bac table: %w", err)
	}
	if err := db.Migrator().CreateTable(&entity.MicroplasticMeasurement{}); err != nil {
		return fmt.Errorf("failed to create microplastic_measurement table: %w", err)
	}
	return nil
}
