package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"inventory-management/initializers"
	"inventory-management/models"
	"inventory-management/repository"
	"time"
)

type RestockService struct {
	ItemRepo    *repository.ItemRepository
	RestockRepo *repository.RestockRepository
	Redis       *redis.Client
}

func NewRestockService(itemRepo *repository.ItemRepository, restockRepo *repository.RestockRepository) *RestockService {
	return &RestockService{
		ItemRepo:    itemRepo,
		RestockRepo: restockRepo,
		Redis:       initializers.RedisClient,
	}
}

func (s *RestockService) RestockItem(itemID int, amount int) (*models.Item, error) {
	// Validate restock amount
	if amount < 10 || amount > 1000 {
		return nil, errors.New("restock amount must be between 10 and 1000")
	}

	ctx := context.Background()
	key := fmt.Sprintf("restock:rate_limit:item:%d", itemID)
	count, err := s.Redis.Incr(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to increment rate limit: %v", err)
	}

	if count == 1 {
		err = s.Redis.Expire(ctx, key, 24*time.Hour).Err()
		if err != nil {
			return nil, fmt.Errorf("failed to set TTL: %v", err)
		}
	}

	if count > 3 {
		return nil, errors.New("rate limit exceeded: max 3 restocks per 24 hours")
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

func (s *RestockService) GetRestockHistory(itemID int) ([]models.RestockHistory, error) {
	return s.RestockRepo.FindByItemID(itemID)
}
