package persistence

import (
	"backend/cmd/api/domain/model"
	"backend/cmd/api/domain/repository"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type itemPersistence struct {
	Conn *sqlx.DB
}

func NewItemPersistence(conn *sqlx.DB) repository.ItemRepository {
	return &itemPersistence{Conn: conn}
}

func (ip *itemPersistence) FindAll() (items []*model.Item, err error) {
	db := ip.Conn

	err = db.Select(&items, "SELECT * FROM items")
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (ip *itemPersistence) Create(item *model.Item) (err error) {
	db := ip.Conn
	now := time.Now()

	_, err = db.Exec(
		"INSERT INTO items (name, status, created_at, updated_at) values (?, ?, ?, ?)",
		item.Name, item.Status, now, now)
	log.Print(item.Status)
	if err != nil {
		return err
	}

	return nil
}

func (ip *itemPersistence) Update(item *model.Item) (err error) {
	db := ip.Conn
	now := time.Now()

	_, err = db.Exec(
		"UPDATE items SET name = ?, status = ?, updated_at = ? WHERE id = ?",
		item.Name, item.Status, now, item.Id)

	return nil
}
