package adapter

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ZeroLogger struct {
	sublogger zerolog.Logger
}

func NewZeroLogger() *ZeroLogger {
	return &ZeroLogger{
		sublogger: zerolog.New(os.Stdout),
	}
}

func (z *ZeroLogger) Log(msg string, keyval ...interface{}) {
	if len(keyval)%2 == 1 {
		panic("expected even number")
	}
	logger := z.sublogger.Log()
	for i := 0; i < len(keyval); i += 2 {
		key := keyval[i].(string)
		logger = logger.Interface(key, keyval[i+1])
	}
	logger.Msg(msg)
}

func (z *ZeroLogger) With(keyval ...interface{}) Logger {
	ctx := zlog.With()
	for i := 0; i < len(keyval); i += 2 {
		ctx = ctx.Interface(keyval[i].(string), keyval[i+1])
	}
	return &ZeroLogger{
		sublogger: ctx.Logger(),
	}
}
func (z *ZeroLogger) log(msg string, logger *zerolog.Event, keyval ...interface{}) {
	for i := 0; i < len(keyval); i += 2 {
		logger = logger.Interface(keyval[i].(string), keyval[i+1])
	}
	logger.Msg(msg)
}
func (z *ZeroLogger) Debug(msg string, keyval ...interface{}) {
	z.log(msg, z.sublogger.Debug(), keyval...)
}
func (z *ZeroLogger) Info(msg string, keyval ...interface{}) {
	z.log(msg, z.sublogger.Info(), keyval...)
}
func (z *ZeroLogger) Error(msg string, keyval ...interface{}) {
	z.log(msg, z.sublogger.Error(), keyval...)
}
