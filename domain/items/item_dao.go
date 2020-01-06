package items

import (
	"errors"
	"fmt"
	"github.com/IkezawaYuki/videostore_items-api/clients/elasticsearch"
	"github.com/IkezawaYuki/videostore_utils-go/rest_errors"
	"strings"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.ID = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	result, err := elasticsearch.Client.Get(indexItems, typeItem, i.ID)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.ID))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.ID), errors.New("database error"))
	}
	fmt.Println(result.Source)

	return nil
}
