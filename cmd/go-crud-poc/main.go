package main

import (
	"log"
	"net/http"

	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/handler"
	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/repository"
	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/service"
)

func main() {
	itemsRepo := repository.NewInMemoryItemsRepository()
	itemsSvc := service.NewItemsService(itemsRepo)
	itemsHandler := handler.NewItemsHandler(itemsSvc)
	http.HandleFunc("/items", itemsHandler.Handle)
	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
