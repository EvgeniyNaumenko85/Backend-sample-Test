package service

import (
	"BST/models"
	"BST/pkg/repository"
	"context"
	"github.com/sirupsen/logrus"
)

type MessageService struct {
	repo repository.Messages
}

func NewMessageService(repo repository.Messages) *MessageService {
	return &MessageService{repo: repo}
}

func (ms *MessageService) AddMessage(ctx context.Context, t models.Message) (id int, err error) {
	id, err = ms.repo.AddMessage(ctx, t)
	if err != nil {
		logrus.Println("app AddMessage", err)
	}
	return
}

func (ms *MessageService) DeleteMessage(ctx context.Context, id, userID int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}
	err = ms.repo.DeleteMessage(ctx, id, userID)
	if err != nil {
		logrus.Println("app DeleteMessage", err)
	}
	return
}

func (ms *MessageService) UpdateMessage(ctx context.Context, id, userID int, t models.Message) (err error) {
	err = ms.repo.UpdateMessage(ctx, id, userID, t)
	if err != nil {
		logrus.Println("app UpdateMessage", err)
	}
	return
}

func (ms *MessageService) GetAllMessages(ctx context.Context, PageID, userID int) (messages []models.Message, err error) {
	messages, err = ms.repo.GetAllMessages(ctx, PageID, userID)
	if err != nil && err != models.ErrNoRows {
		logrus.Println("app GetAllMessages", err)
	}
	return
}
