package repository

import (
	"context"

	"BST/models"
)

type Users interface {
	//Close()
	AddUser(ctx context.Context, u models.User) (err error)
	AuthenticateUser(ctx context.Context, u models.User) (id int, err error)
	GetUser(ctx context.Context, id int) (user models.User, err error)
	GetAllUsers(ctx context.Context) (users []models.User, err error)
	DeleteUser(ctx context.Context, id int) (err error)
	UpdateUser(ctx context.Context, u models.User) (err error)
}

type Messages interface {
	//Close()
	AddMessage(ctx context.Context, t models.Message) (id int, err error)
	DeleteMessage(ctx context.Context, id, userID int) (err error)
	UpdateMessage(ctx context.Context, id, userID int, t models.Message) (err error)
	GetAllMessages(ctx context.Context, PageID, userID int) (messages []models.Message, err error)
}

type Repository struct {
	Users
	Messages
}

func NewRepository(cfg string) *Repository {
	userDB := newConnection(cfg)
	messageDB := newConnection(cfg)
	return &Repository{
		Users:    NewUserDB(userDB),
		Messages: NewMessageDB(messageDB),
	}
}

//func (r *Repository) Close() {
//	r.Users.Close()
//	r.Messages.Close()
//}
