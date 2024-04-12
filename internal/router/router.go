package router

import (
	"net/http"
	"rcp-api-data/internal/config/security"
	controller_account "rcp-api-data/internal/controller/account"
	carcontroller "rcp-api-data/internal/controller/car_controller"
	collectorcontroller "rcp-api-data/internal/controller/collector_controller"
	"rcp-api-data/internal/middleware"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Router(db *gorm.DB, cfg *security.Environment, sugar *zap.SugaredLogger) {
	prefix := cfg.PREFIX
	tokenMiddleware := middleware.NewTokenMiddleware(cfg, sugar)
	corsMiddleware := middleware.NewCorsMiddleware(cfg)
	accountController := controller_account.NewAccountController(db, cfg)
	carController := carcontroller.NewCarController(db)
	carUserController := carcontroller.NewCarUserController(db)
	collectorController := collectorcontroller.NewCollectorController(db)
	http.Handle(prefix+"/information/me", corsMiddleware.Config(tokenMiddleware.ValidateTokenMiddleware(http.HandlerFunc(accountController.GetController()))))
	http.Handle(prefix+"/data-collector/car", corsMiddleware.Config(tokenMiddleware.ValidateTokenMiddleware(http.HandlerFunc(carController.CarController()))))
	http.Handle(prefix+"/data-collector/car/user", corsMiddleware.Config(tokenMiddleware.ValidateTokenMiddleware(http.HandlerFunc(carUserController.CarUserController()))))
	http.Handle(prefix+"/data-collector", corsMiddleware.Config(tokenMiddleware.ValidateTokenMiddleware(http.HandlerFunc(collectorController.CollectorController()))))

}
