package repository

import (
	"MessageProcessing/models"
	"github.com/jmoiron/sqlx"
)

type Message interface {
	Create(message models.Message) (int, error)
	GetCurMessages() ([]models.Message, error)
	GetCompMessages() ([]models.Message, error)
	UpdateStatus(messageId int) error
}

type Repository struct {
	Message
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Message: NewMessagePostgres(db),
	}
}
