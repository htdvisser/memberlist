package memberlist

import (
	"fmt"
	"net"
)

// Logger interface is the most simple logging interface which is implemented by
// the standard library log package.
type Logger interface {
	Printf(format string, v ...interface{})
}

// LevelLogger interface allows logging with different log levels.
type LevelLogger interface {
	Logger
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

func levelLogger(logger Logger) LevelLogger {
	if levelLogger, ok := logger.(LevelLogger); ok {
		return levelLogger
	}
	return &noLevelLogger{logger}
}

type noLevelLogger struct {
	Logger
}

func (l *noLevelLogger) Debugf(format string, v ...interface{}) {
	l.Printf("[DEBUG] memberlist: "+format, v...)
}
func (l *noLevelLogger) Infof(format string, v ...interface{}) {
	l.Printf("[INFO] memberlist: "+format, v...)
}
func (l *noLevelLogger) Warnf(format string, v ...interface{}) {
	l.Printf("[WARN] memberlist: "+format, v...)
}
func (l *noLevelLogger) Errorf(format string, v ...interface{}) {
	l.Printf("[ERR] memberlist: "+format, v...)
}

func LogAddress(addr net.Addr) string {
	if addr == nil {
		return "from=<unknown address>"
	}

	return fmt.Sprintf("from=%s", addr.String())
}

func LogStringAddress(addr string) string {
	if addr == "" {
		return "from=<unknown address>"
	}

	return fmt.Sprintf("from=%s", addr)
}

func LogConn(conn net.Conn) string {
	if conn == nil {
		return LogAddress(nil)
	}

	return LogAddress(conn.RemoteAddr())
}
