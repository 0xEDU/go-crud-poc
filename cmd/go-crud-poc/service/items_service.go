package service

import (
	"errors"

	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/model"
	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/repository"
)

type ItemsService struct {
	repo repository.ItemsRepository
}

func NewItemsService(repo repository.ItemsRepository) *ItemsService {
	return &ItemsService{
		repo: repo,
	}
}

func (service *ItemsService) DeleteItem(oldItem model.Item) (model.Item, error){
	return service.repo.DeleteItem(oldItem)
}

func (service *ItemsService) GetAllItems() ([]model.Item, error) {
	return service.repo.GetAllItems()
}

func (service *ItemsService) GetItem(id int) (model.Item, error) {
	return service.repo.GetItem(id)
}

func (service *ItemsService) UpdateItem(newItem model.Item) (model.Item, error) {
	if newItem.Name == "" {
		return model.Item{}, errors.New("Item name cannot be empty")
	}
	return service.repo.UpdateItem(newItem)
}

func (service *ItemsService) CreateItem(item model.Item) (model.Item, error) {
	if item.Name == "" {
		return model.Item{}, errors.New("Item name cannot be empty")
	}
	err := service.repo.AddItem(&item)
	if err != nil {
		return model.Item{}, err
	}
	return item, nil
}
