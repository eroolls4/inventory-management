package services

import (
	"errors"
	"inventory-management/models"
	"inventory-management/repository"
	"time"
)

type RestockService struct {
	ItemRepo    *repository.ItemRepository
	RestockRepo *repository.RestockRepository
}

func NewRestockService(itemRepo *repository.ItemRepository, restockRepo *repository.RestockRepository) *RestockService {
	return &RestockService{ItemRepo: itemRepo, RestockRepo: restockRepo}
}

func (s *RestockService) RestockItem(itemID int, amount int) (*models.Item, error) {
	if amount < 10 || amount > 1000 {
		return nil, errors.New("restock amount must be between 10 and 1000")
	}

	item, err := s.ItemRepo.FindByID(itemID)
	if err != nil {
		return nil, err
	}

	item.Quantity += amount
	item.LastRestockAt = time.Now()
	item.UpdatedAt = time.Now()
	if err := s.ItemRepo.Update(&item); err != nil {
		return nil, err
	}

	restock := &models.RestockHistory{
		ItemID:           itemID,
		RestockAmount:    amount,
		RestockTimestamp: time.Now(),
	}
	if err := s.RestockRepo.Create(restock); err != nil {
		return nil, err
	}

	return &item, nil
}
