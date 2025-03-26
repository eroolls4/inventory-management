package repository

import (
	"gorm.io/gorm"
	"inventory-management/models"
)

type InventoryRepository struct {
	DB *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{DB: db}
}

func (r *InventoryRepository) Create(inventory *models.Inventory) error {
	return r.DB.Create(inventory).Error
}

func (r *InventoryRepository) FindByID(id int) (models.Inventory, error) {
	var inventory models.Inventory
	err := r.DB.First(&inventory, id).Error
	return inventory, err
}

func (r *InventoryRepository) Update(inventory *models.Inventory) error {
	return r.DB.Save(inventory).Error
}

func (r *InventoryRepository) Delete(id int) error {
	return r.DB.Delete(&models.Inventory{}, id).Error
}
