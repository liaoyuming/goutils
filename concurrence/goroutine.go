package concurrence

import (
	"fmt"
	"runtime"
)

type PanicLogFunc func(format string, v ...interface{})

var (
	defaultPanicLogFunc PanicLogFunc
)

func init() {
	defaultPanicLogFunc = func(format string, v ...interface{}) {
		fmt.Printf(format, v...)
	}
}

func SetPanicLogFunc(l PanicLogFunc) {
	defaultPanicLogFunc = l
}

func DefaultDeferFunc() {
	if err := recover(); err != nil {
		const size = 64 << 10
		buf := make([]byte, size)
		buf = buf[:runtime.Stack(buf, false)]
		defaultPanicLogFunc("goroutine panic: %s  stack: %s", err, buf)
	}
}

func GoWithRecovery(df func(), f func()) {
	go func() {
		defer df()
		f()
	}()
}

func GoWithDefaultRecovery(f func()) {
	GoWithRecovery(DefaultDeferFunc, f)
}
