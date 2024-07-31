package services

import (
	"encoding/json"
	"module/internal/models"

	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
)

func producerSend(message models.Grades) error {

	// структуру в джейсон
	orderInBytes, err := json.Marshal(message)
	if err != nil {
		log.Error(err)
		return models.ResponseBadRequest()
	}

	// отправка джейсона в кафку
	err = pushOrderToQueue("math_send", orderInBytes)
	if err != nil {
		log.Error("kafka send error")
		return models.ResponseErrorAtServer()
	}

	return models.ResponseGood()
}

func connectProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	return sarama.NewSyncProducer(brokers, config)
}

func pushOrderToQueue(topic string, message []byte) error {

	brokers := []string{kafkaAddress}

	// соединение
	producer, err := connectProducer(brokers)
	if err != nil {
		return err
	}

	defer producer.Close()

	// создание сообщения
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// отправка сообщения
	_, _, err = producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}
