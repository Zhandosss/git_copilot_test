package internal

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectURL  string   `json:"redirect_url"`
	Scopes       []string `json:"scopes"`
	Provider     string   `json:"provider"`
}

func NewConfig() *Config {
	config := &Config{}
	raw, err := os.ReadFile("config.json")
	log.Println("raw:", string(raw))
	if err != nil {
		log.Println("failed to read config file:", err)
		return nil
	}
	err = json.Unmarshal(raw, config)
	log.Println("config:", config)
	if err != nil {
		log.Println("failed to unmarshal config file:", err)
		return nil
	}
	return config
}
