package migration

import (
	"rcp-api-data/internal/entity"

	"gorm.io/gorm"
)

func RunMigration (db *gorm.DB) {
	println("Running migration")
	createEnums(db)
	println("Enums created")
	db.Migrator().DropTable(&entity.Brand{})
	db.Migrator().DropTable(&entity.Car{})
	println("Tables cleaned")
	db.Migrator().CreateTable(&entity.Brand{})
	db.Migrator().CreateTable(&entity.Car{})
	println("Tables created")
}

func createEnums(db *gorm.DB) error {
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