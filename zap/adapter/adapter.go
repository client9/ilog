package adapter

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() *ZapLogger {
	lg, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return &ZapLogger{
		logger: lg.Sugar(),
	}
}

func (zl *ZapLogger) Log(msg string, keyval ...interface{}) {
	zl.Debug(msg, keyval...)
}

func (zl *ZapLogger) With(keyval ...interface{}) Logger {
	return &ZapLogger{
		logger: zl.logger.With(keyval...),
	}
}

func (zl *ZapLogger) Debug(msg string, keyval ...interface{}) {
	zl.logger.Debugw(msg, keyval...)
}
func (zl *ZapLogger) Info(msg string, keyval ...interface{}) {
	zl.logger.Infow(msg, keyval...)
}
func (zl *ZapLogger) Error(msg string, keyval ...interface{}) {
	zl.logger.Errorw(msg, keyval...)
}
