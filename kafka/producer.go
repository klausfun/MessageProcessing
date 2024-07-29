package kafka

import (
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
)

type Producer struct {
	asyncProducer sarama.AsyncProducer
}

var once sync.Once
var instance *Producer

func GetProducer() *Producer {
	once.Do(func() {
		config := sarama.NewConfig()
		config.Producer.Return.Successes = true
		config.Producer.Return.Errors = true

		producer, err := sarama.NewAsyncProducer(viper.GetStringSlice("kafka.brokers"), config)
		if err != nil {
			logrus.Fatalf("Failed to start Sarama producer: %s", err)
		}

		instance = &Producer{asyncProducer: producer}

		go func() {
			for {
				select {
				case msg := <-producer.Successes():
					logrus.Infof("Message published: %v", msg)
				case err := <-producer.Errors():
					logrus.Errorf("Failed to publish message: %v", err)
				}
			}
		}()
	})

	return instance
}

func (p *Producer) SendMessage(topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	p.asyncProducer.Input() <- msg
	return nil
}
