package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/model"
	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/service"
	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/view"
)

type ItemsHandler struct {
	svc *service.ItemsService
}

func NewItemsHandler(svc *service.ItemsService) *ItemsHandler {
	return &ItemsHandler{
		svc: svc,
	}
}

func (handler *ItemsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Println("Get items")
		handler.getItems(w)
	case "POST":
		log.Println("Create item")
		handler.createItem(w, r)
	case "PUT":
		log.Println("Update item")
		handler.updateItem(w, r)
	case "DELETE":
		log.Println("Delete item")
		handler.deleteItem(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (handler *ItemsHandler) deleteItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	deletedItem, err := handler.svc.DeleteItem(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.JSON(w, http.StatusOK, deletedItem)
}

func (handler *ItemsHandler) updateItem(w http.ResponseWriter, r *http.Request) {
	var newItem model.Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedItem, err := handler.svc.UpdateItem(newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.JSON(w, http.StatusOK, updatedItem)
}

func (handler *ItemsHandler) createItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdItem, err := handler.svc.CreateItem(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.JSON(w, http.StatusCreated, createdItem)
}

func (handler *ItemsHandler) getItems(w http.ResponseWriter) {
	items, err := handler.svc.GetAllItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.JSON(w, http.StatusOK, items)
}
