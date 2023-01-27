package gerror

type LogLevel string

const (
	LevelUnknown LogLevel = ""
	LevelNever   LogLevel = "never"
	LevelDebug   LogLevel = "debug"
	LevelInfo    LogLevel = "info"
	LevelWarn    LogLevel = "warn"
	LevelError   LogLevel = "error"
)

// LogInterface is a subset of github.com/grafana/grafana/pkg/infra/log.Logger
// to avoid having to depend on other packages in the module so that
// there's no risk of circular dependencies.
type LogInterface interface {
	Debug(msg string, ctx ...any)
	Info(msg string, ctx ...any)
	Warn(msg string, ctx ...any)
	Error(msg string, ctx ...any)
}

func (l LogLevel) LogFunc(logger LogInterface) func(msg string, ctx ...any) {
	switch l {
	case LevelNever:
		return func(_ string, _ ...any) {}
	case LevelDebug:
		return logger.Debug
	case LevelInfo:
		return logger.Info
	case LevelWarn:
		return logger.Warn
	default: // LevelUnknown and LevelError
		return logger.Error
	}
}
