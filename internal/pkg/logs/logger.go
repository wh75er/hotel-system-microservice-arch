package logs

import "github.com/sirupsen/logrus"

type LoggerInterface interface {
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type LogrusInstance struct {
	logger *logrus.Logger
}

func NewLogrus() LoggerInterface {
	return &LogrusInstance{
		logrus.New(),
	}
}

func (l *LogrusInstance) Info(v ...interface{}) {
	l.logger.Info(v...)
}

func (l *LogrusInstance) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

func (l *LogrusInstance) Error(v ...interface{}) {
	l.logger.Error(v...)
}

func (l *LogrusInstance) Errorf(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
}

func (l *LogrusInstance) Warn(v ...interface{}) {
	l.logger.Warn(v...)
}

func (l *LogrusInstance) Warnf(format string, v ...interface{}) {
	l.logger.Warnf(format, v...)
}

func (l *LogrusInstance) Fatal(v ...interface{}) {
	l.logger.Fatal(v...)
}

func (l *LogrusInstance) Fatalf(format string, v ...interface{}) {
	l.logger.Fatalf(format, v...)
}
