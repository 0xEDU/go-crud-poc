package repository

import (
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

func AddItem(item *model.Item) error {
	mu.Lock()
	defer mu.Unlock()
	item.ID = len(items) + 1
	items = append(items, *item)
	return nil
}
