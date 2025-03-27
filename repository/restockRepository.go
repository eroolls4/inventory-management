package repository

import (
	"gorm.io/gorm"
	"inventory-management/models"
)

type RestockRepository struct {
	DB *gorm.DB
}

func NewRestockRepository(db *gorm.DB) *RestockRepository {
	return &RestockRepository{DB: db}
}

func (r *RestockRepository) Create(restock *models.RestockHistory) error {
	return r.DB.Create(restock).Error
}

func (r *RestockRepository) FindByID(id int) (models.RestockHistory, error) {
	var restock models.RestockHistory
	err := r.DB.First(&restock, id).Error
	return restock, err
}

func (r *RestockRepository) FindByItemID(itemID int) ([]models.RestockHistory, error) {
	var history []models.RestockHistory
	err := r.DB.Where("item_id = ?", itemID).Find(&history).Error
	return history, err
}
