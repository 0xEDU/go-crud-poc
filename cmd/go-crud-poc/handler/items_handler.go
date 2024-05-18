package handler

import (
	"log"
	"net/http"

	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/service"
	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/view"
)

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Println("Get items")
		getItems(w)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getItems(w http.ResponseWriter) {
	items, err := service.GetAllItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.JSON(w, http.StatusOK, items)
}
