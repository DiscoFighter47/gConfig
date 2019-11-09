package gconfig

import (
	"log"
	"os"
)

// Logger ...
type Logger interface {
	Fatal(v ...interface{})
}

var logger Logger = log.New(os.Stdout, "", log.LstdFlags)

// SetLogger ...
func SetLogger(l Logger) {
	logger = l
}
