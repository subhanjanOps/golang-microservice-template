package appLogger

import (
	customError "app-server-gateway-service/pkg/custom_error"
	"errors"
	"go.uber.org/zap"
	"os"
)

type IServerLogger interface {
	GetLogger() *zap.Logger
	InfoLog()
	WarningLog()
	ErrorLog()
}

type ServerLogger struct {
	logger *zap.Logger
}
type ApiLogStruct struct {
	method string
	status int
	url    string
}

func NewServerLogger() *ServerLogger {
	var logger *zap.Logger
	if os.Getenv("APP_ENV") == "PROD" {
		logger = productionConfig()
	}
	if os.Getenv("APP_ENV") == "DEV" {
		logger = developmentConfig()
	} else {
		customError.CheckError(errors.New("APP_ENV must be set and valid at environment"))
	}
	return &ServerLogger{
		logger: logger,
	}
}

func productionConfig() *zap.Logger {
	//conf := zap.NewProductionConfig()

	lg, err := zap.NewProduction()
	customError.CheckError(err)
	return lg
}
func developmentConfig() *zap.Logger {
	lg, err := zap.NewDevelopment()
	customError.CheckError(err)
	return lg
}
