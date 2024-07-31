package services

import (
	"encoding/json"
	"module/internal/database"
	"module/internal/models"

	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
)

var kafkaAddress string

func KafkaUrlAdd(address string) {
	kafkaAddress = address
}

func KafkaConsumer() {
	topic := "math_send"

	// создание консьюмера
	worker, err := ConnectConsumer([]string{kafkaAddress})
	if err != nil {
		panic(err)
	}
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	// горутина для запуска консьюмера
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Error(err)
			case msg := <-consumer.Messages():
				log.Info("kafka accepted message")

				// Обновляем статистику что пришло сообщение.
				database.StatisticAddByMetric("AllRecords")

				// парсим в джейсон
				sendedMessage := &models.Grades{}
				err := json.Unmarshal(msg.Value, sendedMessage)
				if err != nil {
					log.Error("error at parse")
					break
				}

				if sendedMessage.Mathematics > 100 || sendedMessage.Mathematics < 0 || sendedMessage.RusLanguage > 100 || sendedMessage.RusLanguage < 0 {
					log.Info("invalid data, doesn't verify")
					break
				}

				// Обновляем статистику что сообщение верифицировано.
				database.StatisticAddByMetric("Verified")
			}
		}
	}()

	<-doneCh

}

func ConnectConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	return sarama.NewConsumer(brokers, config)
}
