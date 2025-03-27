package services

import (
	"inventory-management/models"
	"inventory-management/repository"
	"time"
)

type InventoryService struct {
	Repo *repository.InventoryRepository
}

func NewInventoryService(repo *repository.InventoryRepository) *InventoryService {
	return &InventoryService{Repo: repo}
}

func (s *InventoryService) Create(inventory *models.Inventory) (*models.Inventory, error) {
	inventory.CreatedAt = time.Now()
	inventory.UpdatedAt = time.Now()
	err := s.Repo.Create(inventory)
	return inventory, err
}

func (s *InventoryService) GetByID(id int) (models.Inventory, error) {
	return s.Repo.FindByID(id)
}

func (s *InventoryService) Update(id int, updatedInventory *models.Inventory) (*models.Inventory, error) {
	inventory, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	inventory.Name = updatedInventory.Name
	inventory.Description = updatedInventory.Description
	inventory.UpdatedAt = time.Now()
	err = s.Repo.Update(&inventory)
	return &inventory, err
}

func (s *InventoryService) Delete(id int) error {
	return s.Repo.Delete(id)
}

func (s *InventoryService) GetAll() ([]models.Inventory, error) {
	return s.Repo.FindAll()
}
