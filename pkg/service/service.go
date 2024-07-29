package service

import (
	"MessageProcessing/models"
	"MessageProcessing/pkg/repository"
)

type Message interface {
	Create(message models.Message) (int, error)
	GetCurMessages() ([]models.Message, error)
	GetCompMessages() ([]models.Message, error)
	sendToKafka(message models.Message)
	ScanAndResend()
}

type Service struct {
	Message
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Message: NewMessageService(repos.Message),
	}
}
