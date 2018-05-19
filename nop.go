package ilog

type NopLogger struct{}

func (n NopLogger) With(keyval ...interface{}) Logger        { return &n }
func (n NopLogger) Debug(msg string, keyvals ...interface{}) {}
func (n NopLogger) Info(msg string, keyvals ...interface{})  {}
func (n NopLogger) Error(msg string, keyvals ...interface{}) {}
