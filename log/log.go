package log

import (
	"bytes"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const flag = log.Ldate | log.Ltime

// Level describes log level
type Level int

// Log levels
const (
	LevelTrace Level = iota
	LevelDebug
	LevelError
	LevelFatal
)

// Different Level Logger
var (
	Trace = log.New(os.Stdout, "[T] ", flag)
	Debug = log.New(os.Stdout, "[D] ", flag)
	Error = log.New(os.Stdout, "[E] ", flag)
	Fatal = log.New(os.Stderr, "[F] ", flag)
)

// default logLevel is trace
var logLevel Level

// SetLogLevel set log level
func SetLogLevel(level Level) {
	if level >= LevelTrace && level <= LevelFatal {
		logLevel = level
	}
}

// Fatalf mirror log Fatalf
func Fatalf(format string, args ...interface{}) { Fatal.Fatalf(backTrace()+format+"\n", args...) }

// Errorf mirror log Printf
func Errorf(format string, args ...interface{}) {
	if logLevel <= LevelError {
		Error.Printf(backTrace()+format+"\n", args...)
	}
}

// Debugf mirror log Printf
func Debugf(format string, args ...interface{}) {
	if logLevel <= LevelDebug {
		Debug.Printf(backTrace()+format+"\n", args...)
	}
}

// Tracef mirror log Printf
func Tracef(format string, args ...interface{}) {
	if logLevel <= LevelTrace {
		Trace.Printf(backTrace()+format+"\n", args...)
	}
}

func backTrace() string {
	body := bytes.Buffer{}
	body.WriteByte('\n')

	for skip := 2; ; skip++ {
		_, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		file = file[strings.Index(file, "/src/")+5 : len(file)]
		if strings.HasPrefix(file, "runtime/") {
			continue
		}
		body.WriteByte('\t')
		body.WriteString(file)
		body.WriteString("#")
		body.WriteString(strconv.Itoa(line))
		body.WriteByte('\n')
	}

	return body.String()
}
