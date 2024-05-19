package repository

import (
	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/model"
)

type ItemsRepository interface {
	AddItem(item *model.Item) error
	DeleteItem(oldItem model.Item) (model.Item, error)
	GetAllItems() ([]model.Item, error)
	GetItem(id int) (model.Item, error)
	UpdateItem(newItem model.Item) (model.Item, error)
}
