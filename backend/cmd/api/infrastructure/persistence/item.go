package persistence

import (
	"backend/cmd/api/domain/model"
	"backend/cmd/api/domain/repository"
	"github.com/jinzhu/gorm"
	"time"
)

type itemPersistence struct {
	Conn *gorm.DB
}

func NewItemPersistence(conn *gorm.DB) repository.ItemRepository {
	return &itemPersistence{Conn: conn}
}

func (ip *itemPersistence) FindAll() (items []*model.Item, err error) {
	db := ip.Conn

	if err := db.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (ip *itemPersistence) Create(item *model.Item) (err error) {
	db := ip.Conn
	now := time.Now()

	createItem := &model.Item{
		Name:      item.Name,
		Status:    item.Status,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := db.Create(&createItem).Error; err != nil {
		return err
	}

	return nil
}

func (ip *itemPersistence) Update(item *model.Item) (err error) {
	db := ip.Conn
	now := time.Now()

	updateItem := &model.Item{}
	if err := db.First(&updateItem, item.Id).Error; err != nil {
		return err
	}

	updateItem.Name = item.Name
	updateItem.Status = item.Status
	updateItem.UpdatedAt = now

	if err := db.Save(&updateItem).Error; err != nil {
		return err
	}

	return nil
}
