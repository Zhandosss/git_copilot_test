package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	// ClientID is the OAuth client ID used for authentication.
	ClientID string `json:"client_id"`
	// ClientSecret is the OAuth client secret used for authentication.
	ClientSecret string `json:"client_secret"`
	// RedirectURL is the URL to redirect users to after authentication.
	RedirectURL string `json:"redirect_url"`
	// Scopes is a list of OAuth scopes to request.
	Scopes []string `json:"scopes"`
	// Endpoint is the OAuth endpoint to use.
	Provider string `json:"endpoint"`
}

func NewConfig() *Config {
	var config Config
	raw, err := os.ReadFile("config.json")
	if err != nil {
		log.Println("failed to read config file:", err)
		return nil
	}
	err = json.Unmarshal(raw, &config)
	if err != nil {
		log.Println("failed to unmarshal config file:", err)
		return nil
	}
	return &config
}
