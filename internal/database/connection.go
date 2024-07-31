package database

import (
	"fmt"
	"module/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

var GlobalHandler *gorm.DB

// открыть соединение, и поместить его в глобальную переменну.
func OpenConnection(config models.DatabaseConfig) {
	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Pass, config.Host, config.Port, config.Name)

	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		log.Error("database doesn't open")
		panic("database doesn't open")
	}

	GlobalHandler = db
}

// миграция
func StartMigration() {

	GlobalHandler.AutoMigrate(models.Grades{})

	// добавить две записи в статистику
	if !GlobalHandler.Migrator().HasTable(models.Statistic{}) {

		if GlobalHandler.AutoMigrate(models.Statistic{}) == nil {

			verified := models.Statistic{Metrics: "Verified", Value: 0}
			basic := models.Statistic{Metrics: "AllRecords", Value: 0}

			if result := GlobalHandler.Create(&verified); result.Error != nil {
				return
			}
			if result := GlobalHandler.Create(&basic); result.Error != nil {
				return
			}
		}
	}

	log.Info("migration complete")

}

// проверка если есть база данных с нужным именем. если нет, то создать.
func CheckDatabaseCreated(config models.DatabaseConfig) error {

	// открытие соеднение с базой по стандарту
	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Pass, config.Host, config.Port, "postgres")
	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		log.Info("database doesn't open")
		return models.ResponseBadRequest()
	}

	// закрытие бд
	sql, _ := db.DB()
	defer func() {
		_ = sql.Close()
	}()

	// проверка если есть нужная нам база данных
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", config.Name)
	rs := db.Raw(stmt)
	if rs.Error != nil {
		log.Info("error, database dont readed")
		return models.ResponseBadRequest()
	}

	// если нет, то создать
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", config.Name)
		if rs := db.Exec(stmt); rs.Error != nil {
			log.Info("error, dont create a new database")
			return models.ResponseBadRequest()
		}
	}

	return nil
}
