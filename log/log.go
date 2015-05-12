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

// Fatalf mirror log Fatalf
func Fatalf(format string, args ...interface{}) { Fatal.Fatalf(getCallerInfo()+format, args...) }

func getCallerInfo() string {
	b := bytes.NewBuffer(nil)
	_, file, line, ok := runtime.Caller(2)
	if ok {
		b.WriteString(strings.TrimPrefix(file, gopath) + "#" + strconv.Itoa(line) + " ~ ")
	}
	return b.String()
}
