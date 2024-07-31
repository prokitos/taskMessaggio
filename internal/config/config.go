package config

import (
	"module/internal/models"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// получение конфигов из yaml в структуру
func ConfigMustLoad(envName string) *models.Config {

	path := "internal/config/" + envName + ".yaml"
	var cfg models.Config

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("nothing from this path")
	}

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read")
	}

	return &cfg
}
