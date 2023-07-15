package repository

import (
	"BST/models"
	"context"
	"gorm.io/gorm"
)

type MessageDB struct {
	db *gorm.DB
}

func NewMessageDB(db *gorm.DB) *MessageDB {
	return &MessageDB{
		db: db,
	}
}

//func (d *MessageDB) Close() {
//	closeConnection(d.db, "MessageDB")
//}

func (d *MessageDB) AddMessage(ctx context.Context, m models.Message) (id int, err error) {
	err = d.db.Create(&m).Error
	if err == nil {
		id = m.ID
	}
	return
}

func (d *MessageDB) DeleteMessage(ctx context.Context, id, userID int) (err error) {
	err = d.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Message{}).Error
	if err == gorm.ErrRecordNotFound {
		err = models.ErrNoRows
	}
	return
}

func (d *MessageDB) UpdateMessage(ctx context.Context, id, userID int, t models.Message) (err error) {
	err = d.db.Where("id = ? AND user_id = ?", id, userID).Updates(&t).Error
	if err == gorm.ErrRecordNotFound {
		err = models.ErrNoRows
	}
	return
}

func (d *MessageDB) GetAllMessages(ctx context.Context, PageID, userID int) (messages []models.Message, err error) {
	pageSize := 3
	offset := (PageID - 1) * pageSize

	err = d.db.Scopes(d.Pagination(offset, pageSize)).Find(&messages, "user_id = ?", userID).Error
	if err == gorm.ErrRecordNotFound {
		err = models.ErrNoRows
	}
	return
}

func (d *MessageDB) Pagination(offset, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}
}
