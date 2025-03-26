package services

import (
	"inventory-management/models"
	"inventory-management/repository"
	"time"
)

type ItemService struct {
	Repo *repository.ItemRepository
}

func NewItemService(repo *repository.ItemRepository) *ItemService {
	return &ItemService{Repo: repo}
}

func (s *ItemService) Create(item *models.Item) (*models.Item, error) {
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	err := s.Repo.Create(item)
	return item, err
}

func (s *ItemService) GetByID(id int) (models.Item, error) {
	return s.Repo.FindByID(id)
}

func (s *ItemService) Update(id int, updatedItem *models.Item) (*models.Item, error) {
	item, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	item.Name = updatedItem.Name
	item.Description = updatedItem.Description
	item.Quantity = updatedItem.Quantity
	item.InventoryID = updatedItem.InventoryID
	item.UpdatedAt = time.Now()
	err = s.Repo.Update(&item)
	return &item, err
}

func (s *ItemService) Delete(id int) error {
	return s.Repo.Delete(id)
}
