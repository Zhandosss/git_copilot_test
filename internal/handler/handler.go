package handler

import (
	"git_copilot_test/internal/config"
	"git_copilot_test/internal/endpoint"
	"git_copilot_test/internal/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"os"
	"strings"
)

type Handler struct {
	cfg      *config.Config
	services *service.Service
}

func New(cfg *config.Config, services *service.Service) *Handler {
	return &Handler{
		cfg:      cfg,
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	capsule := r.Group("/capsule")
	{
		capsuleOauth := &oauth2.Config{
			ClientID:     h.cfg.Capsule.ClientId,
			ClientSecret: os.Getenv("CAPSULE_CLIENT_SECRET"),
			RedirectURL:  "http://" + h.cfg.Server.Host + h.cfg.Server.Port + "/capsule/callback",
			Scopes:       strings.Split(h.cfg.Capsule.Scope, " "),
			Endpoint:     endpoint.Capsule,
		}

		capsule.GET("/login", h.capsuleLogin(capsuleOauth))
		capsule.GET("/callback", h.capsuleCallback(capsuleOauth))
	}

	r.GET("/info", h.info())

	return r
}
