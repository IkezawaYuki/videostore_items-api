package services

import (
	"github.com/IkezawaYuki/videostore_items-api/domain/items"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *errors.RestErr)
	Get(string) (*items.Item, *errors.RestErr)
}

type itemsService struct {
}

func (s *itemsService) Create(item items.Item) (*items.Item, *errors.RestErr) {
	return nil, nil
}

func (s *itemsService) Get(ID string) (*items.Item, *errors.RestErr) {
	return nil, nil
}
