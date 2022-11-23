package repository

import "backend/cmd/api/domain/model"

type ItemRepository interface {
	FindAll() (items []*model.Item, err error)
	Create(item *model.Item) (err error)
	Update(item *model.Item) (err error)
}
