package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type CapsuleConfig struct {
	ClientId string `mapstructure:"client_id"`
	Scope    string `mapstructure:"scope"`
}

type Config struct {
	Server  ServerConfig  `mapstructure:"http_server"`
	Capsule CapsuleConfig `mapstructure:"capsule_crm"`
}

func New() *Config {
	viper.SetDefault("port", 8080)
	viper.SetDefault("host", "localhost")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("failed to read config file:", err)
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatal("failed to unmarshal config:", err)
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return config

}
