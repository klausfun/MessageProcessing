package service

import (
	"MessageProcessing/models"
	"MessageProcessing/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Message interface {
	Create(message models.Message) (int, error)
	GetCurMessages() ([]models.Message, error)
	GetCompMessages() ([]models.Message, error)
	SendToKafka(message models.Message)
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
