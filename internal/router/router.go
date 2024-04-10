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
	tokenMiddleware := middleware.NewTokenMiddleware(cfg, sugar)
	corsMiddleware := middleware.NewCorsMiddleware(cfg)
	accountController := controller_account.NewAccountController(db, cfg)
	http.Handle(prefix+"/information/me", corsMiddleware.Config(tokenMiddleware.ValidateTokenMiddleware(http.HandlerFunc(accountController.GetController()))))
}
