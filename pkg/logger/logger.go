package logger

import log "github.com/sirupsen/logrus"

type Logger interface {
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Info(args ...interface{})
}

//type Server struct {
//	config *Config
//}
//
//func New() *Server {
//	return &Server{
//		config: buildMyConfigSomehow(),
//	}
//}

func CreateLogger() *log.Logger {
	logger := log.New()
	//logger.Formatter = &log.JSONFormatter{}

	// Use logrus for standard log output
	// Note that `log` here references stdlib's log
	// Not logrus imported under the name `log`.
	log.SetOutput(logger.Writer())
	return logger
}
