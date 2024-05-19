package repository

import (
	"errors"
	"sync"

	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/model"
)

type InMemoryItemsRepository struct {
	items []model.Item
	mu sync.Mutex
}

func NewInMemoryItemsRepository() *InMemoryItemsRepository {
	return &InMemoryItemsRepository{
		items: []model.Item{},
	}
}

func (repo *InMemoryItemsRepository) GetAllItems() ([]model.Item, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	return repo.items, nil
}

func (repo *InMemoryItemsRepository) GetItem(id int) (model.Item, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	for _, item := range repo.items {
		if item.ID == id {
			return item, nil
		}
	}
	return model.Item{}, errors.New("Item not found.")
}

func (repo *InMemoryItemsRepository) UpdateItem(newItem model.Item) (model.Item, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	for i, item := range repo.items {
		if item.ID == newItem.ID {
			repo.items[i] = newItem
			return repo.items[i], nil
		}
	}
	return model.Item{}, errors.New("Couldn't update item.")
}

func (repo *InMemoryItemsRepository) AddItem(item *model.Item) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	item.ID = len(repo.items) + 1
	repo.items = append(repo.items, *item)
	return nil
}

func (repo *InMemoryItemsRepository) DeleteItem(oldItem model.Item) (model.Item, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	index := -1
	for i, item := range repo.items {
		if item.ID == oldItem.ID {
			index = i
			break
		}
	}
	if index != -1 {
		repo.items = append(repo.items[:index], repo.items[index+1:]...)
		return oldItem, nil
	}
	return model.Item{}, errors.New("Couldn't delete item.")
}
