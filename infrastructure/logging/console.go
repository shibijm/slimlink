package logging

import (
	"fmt"
	"slimlink/core/ports"
	"time"
)

type consoleLogger struct{}

func NewConsoleLogger() ports.Logger {
	return &consoleLogger{}
}

func (*consoleLogger) Log(format string, args ...any) {
	format = fmt.Sprintf("[%s] %s\n", time.Now().Format("2006-01-02T15:04:05"), format)
	fmt.Printf(format, args...)
}

func (l *consoleLogger) LogError(err error, source string) {
	l.Log(fmt.Sprintf("[ERROR] [%s] %s", source, err))
}
