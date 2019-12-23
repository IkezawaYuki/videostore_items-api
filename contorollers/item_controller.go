package contorollers

import (
	"fmt"
	"github.com/IkezawaYuki/videostore_items-api/domain/items"
	"github.com/IkezawaYuki/videostore_items-api/services"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	// TODO

	item := items.Item{
		Seller: 0, //TODO GetCallerID(r)
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		//TODO: Return error json to the caller
	}

	fmt.Println(result)
}

func Get(w http.ResponseWriter, r *http.Request) {
	return
}
