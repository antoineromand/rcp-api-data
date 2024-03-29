package router

import (
	"net/http"
	"rcp-api-data/internal/config/security"
	controller_account "rcp-api-data/internal/controller/account"
	"rcp-api-data/internal/middleware"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Router(db *gorm.DB, cfg *security.Environment, sugar *zap.SugaredLogger) {
	prefix := cfg.PREFIX
	http.Handle(prefix+"/information/me", middleware.CorsMiddleware(middleware.ValidateTokenMiddleware(http.HandlerFunc(controller_account.AccountController(db)), cfg, sugar), cfg))
}
