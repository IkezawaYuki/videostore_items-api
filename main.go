package main

import (
	"github.com/IkezawaYuki/videostore_items-api/contorollers"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	router = mux.NewRouter()
)

func main() {
	router.HandleFunc("/items", contorollers.Create).Methods(http.MethodPost)
}
