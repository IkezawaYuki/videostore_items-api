package contorollers

import (
	"fmt"
	"github.com/IkezawaYuki/videostore_items-api/domain/items"
	"github.com/IkezawaYuki/videostore_items-api/services"
	"net/http"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
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

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	return
}
