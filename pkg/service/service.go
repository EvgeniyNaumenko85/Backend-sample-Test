package service

import (
	"context"

	"BST/models"
	"BST/pkg/repository"
)

type Authorization interface {
	GenerateToken(ctx context.Context, u models.User) (string, error)
	ParseToken(jwtString string) (int, error)
}

type Message interface {
	AddMessage(ctx context.Context, t models.Message) (id int, err error)
	GetAllMessages(ctx context.Context, id, userID int) (messages []models.Message, err error)
	UpdateMessage(ctx context.Context, id, userID int, t models.Message) (err error)
	DeleteMessage(ctx context.Context, id, userID int) (err error)
}

type Users interface {
	GetUser(ctx context.Context, userID int) (user models.User, err error)
	GetAllUsers(ctx context.Context) (users []models.User, err error)
	AddUser(ctx context.Context, u models.User) (err error)
	DeleteUser(ctx context.Context, userID int) (err error)
	UpdateUser(ctx context.Context, u models.User) (err error)
}

type Service struct {
	Authorization
	Users
	Message
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Users),
		Users:         NewUserService(repos.Users),
		Message:       NewMessageService(repos.Messages),
	}
}
