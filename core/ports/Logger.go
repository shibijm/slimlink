package ports

type Logger interface {
	Log(format string, a ...any)
	LogError(err error, source string)
}
