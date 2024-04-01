package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git_copilot_test/internal/config"
	"git_copilot_test/internal/endpoints/capsule"
	"git_copilot_test/internal/model"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

func initOauthConfig(config *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.ClientId,
		ClientSecret: config.ClientSecret,
		RedirectURL:  "http://" + config.Server.Host + config.Server.Port + "/callback",
		Scopes:       strings.Split(config.Scope, " "),
		Endpoint:     capsule.Endpoint,
	}
}

func main() {
	cfg := config.New()
	log.Println("config:", cfg)
	r := gin.Default()

	oauthConf := initOauthConfig(cfg)
	log.Println("oauth config:", oauthConf)

	r.GET("/login", func(c *gin.Context) {
		url := oauthConf.AuthCodeURL("state")
		log.Println("redirect to:", url)
		c.Redirect(http.StatusFound, url)
	})

	r.GET("/callback", func(c *gin.Context) {
		code := c.Query("code")
		token, err := oauthConf.Exchange(c, code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		log.Println("token:", token)
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	r.GET("/info", func(c *gin.Context) {
		systemName := c.Query("systemName")
		if systemName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "systemName is required"})
			return
		}
		message := fmt.Sprintf("For %s how is the API versioned? Get a link for %[1]s where %[1]s publishes API version changelog", systemName)

		reqBody := map[string]interface{}{
			"model": "gpt-3.5-turbo",
			"messages": []map[string]interface{}{
				{
					"role":    "user",
					"content": message,
				},
			},
			"temperature": 0.0,
		}

		jsonBody, err := json.Marshal(reqBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+cfg.OpenAIKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected status code: " + resp.Status})
			return
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var result model.InfoResponse
		err = json.Unmarshal(respBody, &result)
		if err != nil {
			log.Println("failed to unmarshal response:", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": result.Choices[0].Msg.Content})

	})

	log.Fatal(r.Run(cfg.Server.Host + cfg.Server.Port))
}
