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

// Different Level Logger
var (
	Trace = log.New(os.Stdout, "[T] ", flag)
	Debug = log.New(os.Stdout, "[D] ", flag)
	Error = log.New(os.Stdout, "[E] ", flag)
	Fatal = log.New(os.Stderr, "[F] ", flag)
)

var gopath = os.Getenv("GOPATH") + "/src/"
var goroot = os.Getenv("GOROOT") + "/src/"

// Fatalf mirror log Fatalf
func Fatalf(format string, args ...interface{}) { Fatal.Fatalf(backTrace()+format+"\n", args...) }

func backTrace() string {
	body := bytes.NewBufferString("\n")

	for skip := 2; ; skip++ {
		_, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		if file[len(file)-1] == 'c' {
			continue
		}
		body.WriteString("\t" + strings.TrimPrefix(strings.TrimPrefix(file, gopath), goroot) + "#" + strconv.Itoa(line) + "\n")
	}

	return body.String()
}
