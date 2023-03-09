package logging

import (
	"fmt"
	"slimlink/core/ports"
	"time"
)

type ConsoleLogger struct{}

func NewConsoleLogger() ports.Logger {
	return &ConsoleLogger{}
}

func (consoleLogger *ConsoleLogger) Log(format string, args ...any) {
	format = fmt.Sprintf("[%s] %s\n", time.Now().Format("2006-01-02T15:04:05"), format)
	fmt.Printf(format, args...)
}

func (consoleLogger *ConsoleLogger) LogError(err error, source string) {
	consoleLogger.Log(fmt.Sprintf("[ERROR] [%s] %s", source, err))
}
