package thewq

import (
	"fmt"
	"log"
	"os"
	"sync"
)

// Logger logger
type Logger interface {
	Output(int, string) error
}

var logger Logger = log.New(os.Stdout, "INFO: ", log.Lshortfile)
var l sync.RWMutex

func SetLogger(lg Logger) {
	l.Lock()
	logger = lg
	l.Unlock()
}

func Debugf(format string, a ...interface{}) {
	l.RLock()
	o := logger
	l.RUnlock()

	o.Output(2, fmt.Sprintf(format, a...))
}

func Errorf(format string, a ...interface{}) {
	l.RLock()
	o := logger
	l.RUnlock()

	o.Output(4, fmt.Sprintf(format, a...))
}
