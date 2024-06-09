package db

import (
	"log"

	"github.com/anduckhmt146/graphql-api/internal/model"
	"gorm.io/gorm"
)

func autoMigrateSchema(db *gorm.DB) error {
	models := []interface{}{
		&model.User{},
	}
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			log.Printf("Failed to migrate schema: %v\n", err)
		}
	}
	log.Println("Schema migrated successfully!")
	return nil
}
