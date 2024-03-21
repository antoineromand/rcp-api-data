package utils

import (
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func GetLogger() *zap.SugaredLogger {
    if sugar == nil {
        var err error
        logger, err := zap.NewProduction()
		sugar = logger.Sugar()
		defer logger.Sync()
        if err != nil {
            panic(err)
        }
    }
    return sugar
}