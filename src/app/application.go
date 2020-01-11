package app

import (
	"github.com/IkezawaYuki/videostore_items-api/src/clients/elasticsearch"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()

	mapUrls()

	srv := &http.Server{
		Addr:         "127.0.0.1:8081",
		Handler:      router,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 500 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
