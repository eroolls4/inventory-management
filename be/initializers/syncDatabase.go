package initializers

import (
	"inventory-management/models"
	"log"
	"time"
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

	var count int64
	DB.Model(&models.Inventory{}).Count(&count)
	if count == 0 {
		inventories := []models.Inventory{
			{
				Name:        "Retail Store",
				Description: "Main retail location for customer sales",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Central Warehouse",
				Description: "Primary storage and distribution center",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		}

		result := DB.Create(&inventories)
		if result.Error != nil {
			log.Fatalf("Failed to insert initial inventory rows: %v", result.Error)
		}
		log.Printf("Inserted %d initial inventory rows", result.RowsAffected)
	} else {
		log.Println("Inventory rows already exist. Skipping initial insertion.")
	}

	log.Println("Database sync completed successfully")
}
