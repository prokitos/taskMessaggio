package database

import (
	"module/internal/models"

	log "github.com/sirupsen/logrus"
)

func GetStatistics() ([]models.Statistic, error) {

	var finded []models.Statistic

	results := GlobalHandler.Find(&finded)
	if results.Error != nil {
		log.Error("statistic find error")
		return nil, models.ResponseErrorAtServer()
	}

	return finded, models.ResponseGood()

}

func StatisticAddByMetric(statName string) {

	// можео вынести в кафку, чтобы на один вызов чтения был меньше.
	allStat, err := GetStatistics()
	if err.Error() != models.ResponseGood().Error() {
		log.Error("statistic find error")
		return
	}

	// получение конкретной статистики
	var curStat models.Statistic
	for i := 0; i < len(allStat); i++ {
		if allStat[i].Metrics == statName {
			curStat = allStat[i]
			break
		}
	}

	// прибавление на один и обновить в таблице
	curStat.Value++
	GlobalHandler.Model(models.Statistic{}).Where("metrics = ?", curStat.Metrics).Updates(curStat)

}
