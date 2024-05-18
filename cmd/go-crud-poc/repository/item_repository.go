package repository

import (
	"errors"
	"sync"

	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/model"
)

var (
	items = []model.Item{}
	mu sync.Mutex
)

func GetAllItems() ([]model.Item, error) {
	mu.Lock()
	defer mu.Unlock()
	return items, nil
}

func GetItem(id int) (model.Item, error) {
	mu.Lock()
	defer mu.Unlock()
	for _, item := range items {
		if item.ID == id {
			return item, nil
		}
	}
	return model.Item{}, errors.New("Item not found.")
}

func UpdateItem(newItem model.Item) (model.Item, error) {
	mu.Lock()
	defer mu.Unlock()
	for i, item := range items {
		if item.ID == newItem.ID {
			items[i] = newItem
			return items[i], nil
		}
	}
	return model.Item{}, errors.New("Couldn't update item.")
}

func AddItem(item *model.Item) error {
	mu.Lock()
	defer mu.Unlock()
	item.ID = len(items) + 1
	items = append(items, *item)
	return nil
}
