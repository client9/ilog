package adapter

import (
	"os"

	klog "github.com/go-kit/kit/log"
	kloglevel "github.com/go-kit/kit/log/level"
)

type KitLogger struct {
	logger klog.Logger
}

func NewKitLogger() *KitLogger {
	return &KitLogger{
		logger: klog.NewLogfmtLogger(os.Stdout),
	}
}
func (k *KitLogger) With(keyval ...interface{}) Logger {
	return &KitLogger{
		logger: klog.With(k.logger, keyval...),
	}
}
func (k *KitLogger) Debug(msg string, keyval ...interface{}) {
	kloglevel.Debug(k.logger).Log(append([]interface{}{"msg", msg}, keyval...)...)
}
func (k *KitLogger) Info(msg string, keyval ...interface{}) {
	kloglevel.Info(k.logger).Log(append([]interface{}{"msg", msg}, keyval)...)
}
func (k *KitLogger) Error(msg string, keyval ...interface{}) {
	kloglevel.Error(k.logger).Log(append([]interface{}{"msg", msg}, keyval)...)
}
