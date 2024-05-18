package service

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
