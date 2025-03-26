package initializers

import (
	"inventory-management/models"
	"log"
)

func SyncDatabase() {

	if DB == nil {
		log.Fatal("Database connection is nil. Make sure ConnectToDb() is called successfully.")
		return
	}

	err := DB.Migrator().AutoMigrate(&models.User{},
		&models.Inventory{},
		&models.Item{},
		&models.RestockHistory{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}

	log.Println("Database sync completed successfully")
}
