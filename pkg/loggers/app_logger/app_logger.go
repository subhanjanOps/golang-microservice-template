package appLogger

import "go.uber.org/zap"

func (s *ServerLogger) GetLogger() *zap.Logger {
	return s.logger
}

func (s *ServerLogger) InfoLog(msg string, args ApiLogStruct) {
	//s.logger.Info(
	//	msg,
	//	zap.String("method", args.method),
	//	zap.Int("status", args.status),
	//	zap.String("url", args.url),
	//)
}

func (s *ServerLogger) WarningLog() {}

func (s *ServerLogger) ErrorLog() {}
