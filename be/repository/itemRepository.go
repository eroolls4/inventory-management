package repository

import (
	"gorm.io/gorm"
	"inventory-management/models"
)

type ItemRepository struct {
	DB *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{DB: db}
}

func (r *ItemRepository) Create(item *models.Item) error {
	return r.DB.Create(item).Error
}

func (r *ItemRepository) FindByID(id int) (models.Item, error) {
	var item models.Item
	err := r.DB.First(&item, id).Error
	return item, err
}

func (r *ItemRepository) Update(item *models.Item) error {
	return r.DB.Save(item).Error
}

func (r *ItemRepository) Delete(id int) error {
	return r.DB.Delete(&models.Item{}, id).Error
}

func (r *ItemRepository) FindAll() ([]models.Item, error) {
	var items []models.Item
	err := r.DB.Preload("Inventory").Find(&items).Error
	return items, err
}

func (r *ItemRepository) FindLowQuantityItems(threshold int) ([]models.Item, error) {
    var items []models.Item
    err := r.DB.Where("quantity < ?", threshold).Find(&items).Error
    return items, err
}