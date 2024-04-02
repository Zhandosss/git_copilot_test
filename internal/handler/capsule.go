package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func (h *Handler) capsuleLogin(oauth *oauth2.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		url := oauth.AuthCodeURL("state")
		log.Println("redirect to:", url)
		c.Redirect(http.StatusFound, url)
	}
}

func (h *Handler) capsuleCallback(oauth *oauth2.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")
		token, err := oauth.Exchange(c, code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		log.Println("token:", token)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
