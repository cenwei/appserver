package log

import (
	"bytes"
	"log"
	"testing"
)

type fakeWriter struct {
	bytes.Buffer
}

func TestLog(t *testing.T) {
	data := "logger_test\n" // fmt.Printf appends a new line
	f := &fakeWriter{}
	Trace = log.New(f, "", 0)
	Trace.Printf(data)

	if f.String() != data {
		t.Errorf("data is %s but should be %s", f.String(), data)
	}

	f.Reset()
	SetLogLevel(LevelError)
	Tracef("anything, because log level is higher than trace")
	if f.String() != "" {
		t.Errorf("data should be empty string")
	}
}
