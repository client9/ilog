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
	out := "level=" + quoted(level)
	for i := 0; i < len(s.prefix); i += 2 {
		key := s.prefix[i]
		if key == "level" {
			continue
		}
		out += " " + quoted(key) + "=" + quoted(s.prefix[i+1])
	}
	for i := 0; i < len(keyval); i += 2 {
		key := keyval[i]
		if key == "level" {
			continue
		}
		out += " " + quoted(key) + "=" + quoted(keyval[i+1])
	}
	if msg != "" {
		out += " msg=" + quoted(msg)
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

func quoted(value interface{}) string {
	val := fmt.Sprintf("%v", value)
	if !shouldQuote(val) {
		return val
	}
	return fmt.Sprintf("%q", val)
}

func shouldQuote(value string) bool {
	if len(value) == 0 {
		return true
	}
	for i := 0; i < len(value); i++ {
		r := value[i]
		if r <= ' ' || r == '=' || r == '"' || r == '\\' {
			return true
		}
	}
	return false
}
