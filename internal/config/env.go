package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	MercadoPagoKey string `mapstructure:"MERCADOPAGO_KEY"`
	SMTPKey        string `mapstructure:"SMTP_KEY"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("config file not found, using environment variables")
		} else {
			return
		}
	}

	err = viper.Unmarshal(&config)
	return
}
