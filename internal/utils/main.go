package utils

import (
	"go.uber.org/zap"
)

func ErrHandler(err error) {
	if err != nil {
		Logger.Errorf("Error %v", err)
		panic(err)
	}
}


// Returns zap Loggered logger to handle log messages
func getLogger() *zap.Logger {
	logger, _ := zap.NewProduction()

	defer logger.Sync()
	return logger
}

var Logger *zap.SugaredLogger = getLogger().Sugar()
