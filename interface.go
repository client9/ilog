package ilog

type Logger interface {
	With(keyval ...interface{}) Logger
	Debug(msg string, keyvals ...interface{})
	Info(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})
}

type NopLogger struct{}

func (n NopLogger) With(keyval ...interface{}) Logger        { return &n }
func (n NopLogger) Debug(msg string, keyvals ...interface{}) {}
func (n NopLogger) Info(msg string, keyvals ...interface{})  {}
func (n NopLogger) Error(msg string, keyvals ...interface{}) {}

/*
func main() {
	log := StandardLog{}
	//log := NewZeroLogger()
	//log := NewApexLogger()
	//log := NewKitLogger()
	//log := NewZapLogger()
	log.Info("foo", "bar", "hungry", true)

	log.With("foo", "bar").Debug("hungry", true)

	news := log.With("foo", "bar")
	news.Debug("hungry", true)
	log.Info("new", "line")

	log.Info("whatever", "yes")
}
*/
