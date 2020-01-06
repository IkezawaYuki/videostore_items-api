package app

import (
	"github.com/IkezawaYuki/videostore_items-api/contorollers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/ping", contorollers.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/items", contorollers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", contorollers.ItemsController.Get).Methods(http.MethodGet)
}
