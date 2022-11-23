package usecase

import (
	"backend/cmd/api/domain/model"
	"backend/cmd/api/domain/repository"
	"errors"
)

type ItemUseCase interface {
	FindAll() (items []*model.Item, err error)
	Create(status int, name string) (err error)
	Update(id, status int, name string) (err error)
}

type itemUseCase struct {
	itemRepository repository.ItemRepository
}

func NewItemUseCase(ir repository.ItemRepository) ItemUseCase {
	return &itemUseCase{
		itemRepository: ir,
	}
}

func (iu itemUseCase) FindAll() (items []*model.Item, err error) {
	items, err = iu.itemRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (iu itemUseCase) Create(status int, name string) (err error) {
	item := &model.Item{
		Name:   name,
		Status: status,
	}

	if err := validate(item); err != nil {
		return err
	}

	err = iu.itemRepository.Create(item)
	if err != nil {
		return err
	}
	return nil
}

func (iu itemUseCase) Update(id, status int, name string) (err error) {
	item := &model.Item{
		Id:     id,
		Name:   name,
		Status: status,
	}

	if err := validate(item); err != nil {
		return err
	}

	err = iu.itemRepository.Update(item)
	if err != nil {
		return err
	}
	return nil
}

func validate(item *model.Item) error {
	if len(item.Name) >= 200 {
		return errors.New("タスク名は200文字未満で書いてください。")
	}
	return nil
}
