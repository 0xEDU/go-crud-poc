package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/model"
	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/service"
	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/view"
)

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Println("Get items")
		getItems(w)
	case "POST":
		log.Println("Create item")
		createItem(w, r)
	case "PUT":
		log.Println("Update item")
		updateItem(w, r)
	case "DELETE":
		log.Println("Delete item")
		deleteItem(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	deletedItem, err := service.DeleteItem(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.JSON(w, http.StatusOK, deletedItem)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	var newItem model.Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedItem, err := service.UpdateItem(newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.JSON(w, http.StatusOK, updatedItem)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdItem, err := service.CreateItem(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.JSON(w, http.StatusCreated, createdItem)
}

func getItems(w http.ResponseWriter) {
	items, err := service.GetAllItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.JSON(w, http.StatusOK, items)
}
