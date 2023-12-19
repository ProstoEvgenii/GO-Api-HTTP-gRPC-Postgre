package config

import "github.com/spf13/viper"

func LoadConfig() error {
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Установка значений по умолчанию, если нужно
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	// ... (другие параметры конфигурации)

	return nil
}
