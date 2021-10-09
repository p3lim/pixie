package log

import "fmt"

type Level int

const (
	ERROR = iota
	WARNING
	INFO
	DEBUG
)

var logLevel Level = INFO

func SetLevel(level Level) {
	logLevel = level
}

func GetLevel(str string) (Level, error) {
	switch str {
	case "ERROR":
		return ERROR, nil
	case "WARNING":
		return WARNING, nil
	case "INFO":
		return INFO, nil
	case "DEBUG":
		return DEBUG, nil
	default:
		return 0, fmt.Errorf("invalid verbosity level %s", str)
	}
}
