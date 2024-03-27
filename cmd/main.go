package main

import (
	"git_copilot_test/internal"
	"git_copilot_test/internal/endpoints/capsule"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

func initOauthConfig(config *internal.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectURL,
		Scopes:       config.Scopes,
		Endpoint:     capsule.Endpoint,
	}
}

func main() {
	config := internal.NewConfig()
	log.Println("config:", config)
	if config == nil {
		log.Fatal("failed to load config")
	}
	oauthConf := initOauthConfig(config)

	r := gin.Default()

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
	log.Fatal(r.Run(":8080"))
}
