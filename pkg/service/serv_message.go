package service

import (
	"MessageProcessing/kafka"
	"MessageProcessing/models"
	"MessageProcessing/pkg/repository"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
)

type MessageService struct {
	repo repository.Message
}

func NewMessageService(repo repository.Message) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) Create(message models.Message) (int, error) {
	id, err := s.repo.Create(message)
	if err != nil {
		return 0, err
	}

	go s.sendToKafka(message)

	return id, nil
}

func (s *MessageService) GetCurMessages() ([]models.Message, error) {
	return s.repo.GetCurMessages()
}

func (s *MessageService) GetCompMessages() ([]models.Message, error) {
	return s.repo.GetCompMessages()
}

func (s *MessageService) sendToKafka(message models.Message) {
	producer := kafka.GetProducer()
	messageData, err := json.Marshal(message)
	if err != nil {
		logrus.Errorf("Failed to marshal message: %v", err)
		return
	}

	err = producer.SendMessage(viper.GetString("kafka.topic"), messageData)
	if err != nil {
		logrus.Errorf("Failed to send message to Kafka: %v", err)
		return
	}

	err = s.repo.UpdateStatus(message.Id)
	if err != nil {
		logrus.Errorf("Failed to update message status: %v", err)
	}
}

func (s *MessageService) ScanAndResend() {
	messages, err := s.repo.GetCurMessages()
	if err != nil {
		logrus.Errorf("Failed to get current messages: %v", err)
		return
	}

	var wg sync.WaitGroup
	for _, message := range messages {
		wg.Add(1)
		go func(msg models.Message) {
			defer wg.Done()
			s.sendToKafka(msg)
		}(message)
	}
	wg.Wait()
}
