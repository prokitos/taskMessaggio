package main

import (
	"fmt"
	"module/internal/app"
	"module/internal/config"
	"module/internal/database"
	"module/internal/services"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func main() {

	// установка логов. установка чтобы показывать логи debug уровня
	log.SetLevel(log.DebugLevel)
	log.Info("the server is starting")

	// получение конфигов
	cfg := config.ConfigMustLoad("render")

	fmt.Println(cfg.Server.Port)

	// проверка что есть бд, или его создание
	err := database.CheckDatabaseCreated(cfg.Database)
	if err != nil {
		return
	}

	// миграция и подключение к бд.
	database.OpenConnection(cfg.Database)
	database.StartMigration()

	// запуск кафка консьюмера в горутине
	services.KafkaUrlAdd(cfg.Kafka)
	go services.KafkaConsumer()

	// запуск сервера в горутине, чтобы потом нормально звершать приложение
	var application app.App
	go application.NewServer(cfg.Server.Port)

	// в итоге мы обрабатываем завершение приложения, и если мы закрыаем его как либо, то оно выполняет действие из метода Stop
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Stop() // безопасное выключение сервера

}
