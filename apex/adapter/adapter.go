package main

import (
	"os"

	alog "github.com/apex/log"
	alogtext "github.com/apex/log/handlers/text"
)

type ApexLogger struct {
	entry alog.Interface
}

func NewApexLogger() *ApexLogger {
	return &ApexLogger{
		entry: &alog.Logger{
			Handler: alogtext.New(os.Stderr),
		},
	}
}

func (a *ApexLogger) With(keyval ...interface{}) Logger {
	fields := make(alog.Fields)
	for i := 0; i < len(keyval); i += 2 {
		fields[keyval[i].(string)] = keyval[i+1]
	}
	return &ApexLogger{
		entry: a.entry.WithFields(fields),
	}
}

func (a *ApexLogger) log(msg string, level int, keyval ...interface{}) {
	if len(keyval) > 0 {
		fields := make(alog.Fields)
		for i := 0; i < len(keyval); i += 2 {
			fields[keyval[i].(string)] = keyval[i+1]
		}
		a.entry.WithFields(fields).Debug(msg)
		return
	}
	switch level {
	case 0:
		a.entry.Debug(msg)
	case 1:
		a.entry.Info(msg)
	case 2:
		a.entry.Error(msg)
	default:
		panic("unknown level")
	}
}
func (a *ApexLogger) Debug(msg string, keyval ...interface{}) {
	a.log(msg, 0, keyval...)
}
func (a *ApexLogger) Info(msg string, keyval ...interface{}) {
	a.log(msg, 1, keyval...)
}
func (a *ApexLogger) Error(msg string, keyval ...interface{}) {
	a.log(msg, 2, keyval...)
}
