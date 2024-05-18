package service

import (
	"errors"

	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/model"
	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/repository"
)

func GetAllItems() ([]model.Item, error) {
	return repository.GetAllItems()
}

func GetItem(id int) (model.Item, error) {
	return repository.GetItem(id)
}

func UpdateItem(newItem model.Item) (model.Item, error) {
	if newItem.Name == "" {
		return model.Item{}, errors.New("Item name cannot be empty")
	}
	return repository.UpdateItem(newItem)
}

func CreateItem(item model.Item) (model.Item, error) {
	if item.Name == "" {
		return model.Item{}, errors.New("Item name cannot be empty")
	}
	err := repository.AddItem(&item)
	if err != nil {
		return model.Item{}, err
	}
	return item, nil
}
