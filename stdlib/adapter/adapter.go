package adapter

import (
	"fmt"
	"log"

	"github.com/client9/ilog"
)

type StandardLog struct {
	logger *log.Logger
	level  string
	prefix []interface{}
}

func New(logger *log.Logger) *StandardLog {
	return &StandardLog{
		logger: logger,
	}
}

func (s *StandardLog) With(keyval ...interface{}) ilog.Logger {
	news := StandardLog{
		logger: s.logger,
		level:  s.level,
		prefix: s.prefix[:],
	}
	news.prefix = append(news.prefix, keyval...)
	return &news
}

func (s *StandardLog) log(level string, msg string, keyval ...interface{}) {
	out := ""
	keyval = append([]interface{}{"level", level, "msg", msg}, keyval...)
	for i := 0; i < len(s.prefix); i += 2 {
		out += fmt.Sprintf("%s=%v ", s.prefix[i], s.prefix[i+1])
	}
	for i := 0; i < len(keyval); i += 2 {
		out += fmt.Sprintf("%s=%v ", keyval[i], keyval[i+1])
	}
	s.logger.Print(out)
}

func (s *StandardLog) Debug(msg string, keyval ...interface{}) {
	s.log("debug", msg, keyval...)
}
func (s *StandardLog) Info(msg string, keyval ...interface{}) {
	s.log("info", msg, keyval...)
}
func (s *StandardLog) Error(msg string, keyval ...interface{}) {
	s.log("error", msg, keyval...)
}
