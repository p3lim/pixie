package log

import (
	"fmt"
	"log"
)

func Fatal(v ...interface{}) {
	if logLevel >= ERROR {
		log.Fatalf("[FATAL] %v", v...)
	}
}

func Fatalf(format string, v ...interface{}) {
	if logLevel >= ERROR {
		log.Fatalf("[FATAL] %v", fmt.Sprintf(format, v...))
	}
}

func Error(v ...interface{}) {
	if logLevel >= ERROR {
		log.Printf("[ERROR] %v", v...)
	}
}

func Errorf(format string, v ...interface{}) {
	if logLevel >= ERROR {
		log.Printf("[ERROR] %v", fmt.Sprintf(format, v...))
	}
}

func Warning(v ...interface{}) {
	if logLevel >= WARNING {
		log.Printf("[WARNING] %v", v...)
	}
}

func Warningf(format string, v ...interface{}) {
	if logLevel >= WARNING {
		log.Printf("[WARNING] %v", fmt.Sprintf(format, v...))
	}
}

func Info(v ...interface{}) {
	if logLevel >= INFO {
		log.Printf("[INFO] %v", v...)
	}
}

func Infof(format string, v ...interface{}) {
	if logLevel >= INFO {
		log.Printf("[INFO] %v", fmt.Sprintf(format, v...))
	}
}

func Debug(v ...interface{}) {
	if logLevel >= DEBUG {
		log.Printf("[DEBUG] %v", v...)
	}
}

func Debugf(format string, v ...interface{}) {
	if logLevel >= DEBUG {
		log.Printf("[DEBUG] %v", fmt.Sprintf(format, v...))
	}
}
