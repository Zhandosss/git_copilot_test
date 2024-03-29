package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type Config struct {
	Server       ServerConfig `mapstructure:"http_server"`
	ClientId     string       `mapstructure:"client_id"`
	ClientSecret string       `envconfig:"CLIENT_SECRET"`
	Scope        string       `mapstructure:"scope"`
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

	err = viper.BindEnv("client_secret")

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatal("failed to unmarshal config:", err)
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ClientSecret = os.Getenv("CLIENT_SECRET")

	return config

}
