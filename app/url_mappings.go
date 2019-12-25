package app

import (
	"github.com/IkezawaYuki/videostore_items-api/contorollers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/items", contorollers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/ping", contorollers.PingController.Ping).Methods(http.MethodGet)
}
