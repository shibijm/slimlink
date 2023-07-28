package ports

type Logger interface {
	Log(format string, args ...any)
	LogError(err error, source string)
}
