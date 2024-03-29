package main

import (
	"git_copilot_test/internal/config"
	"git_copilot_test/internal/endpoints/capsule"
	"github.com/gin-gonic/gin"
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

	log.Fatal(r.Run(cfg.Server.Host + cfg.Server.Port))
}
