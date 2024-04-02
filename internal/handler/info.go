package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type InfoResponse struct {
	Answer []string `json:"answers"`
}

func (h *Handler) info() gin.HandlerFunc {
	return func(c *gin.Context) {
		systemName := c.Query("systemName")
		if systemName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "systemName is required"})
			return
		}

		prompts, err := h.services.GeneratePrompts(systemName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ansBody := make(map[string]string)
		for _, prompt := range prompts {
			ans, err := h.services.GetAnsFromOpenAi(systemName, prompt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			ansBody[prompt] = ans
		}
		c.JSON(http.StatusOK, ansBody)
	}
}
