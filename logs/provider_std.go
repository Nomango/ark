package logs

import "fmt"

func NewStdProvider() Provider {
	return &stdProvider{}
}

type stdProvider struct{}

var _ Provider = (*stdProvider)(nil)

type stdColor string

const (
	stdColorForegroundRed       = "\033[31m"
	stdColorForegroundGreen     = "\033[32m"
	stdColorForegroundYellow    = "\033[33m"
	stdColorForegroundBlue      = "\033[34m"
	stdColorForegroundPurple    = "\033[35m"
	stdColorForegroundDarkGreen = "\033[36m"
	stdColorForegroundWhite     = "\033[37m"
	stdColorReset               = "\033[0m"
)

func levelColor(level Level) stdColor {
	switch level {
	case LevelDebug:
		return stdColorForegroundBlue
	case LevelInfo:
		return stdColorForegroundWhite
	case LevelNotice:
		return stdColorForegroundGreen
	case LevelWarn:
		return stdColorForegroundYellow
	case LevelError:
		return stdColorForegroundRed
	}
	return stdColorForegroundWhite
}

func (*stdProvider) Write(level Level, msg string) {
	fmt.Printf("%s[%s] %s%s\n", levelColor(level), level, msg, stdColorReset)
}
