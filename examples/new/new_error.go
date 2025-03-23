package example

import (
	"github.com/br8km/tracerr"
)

func main() {
	if err := foo(); err != nil {
		tracerr.PrintSourceColor(err)
	}
}

func foo() error {
	return bar(0)
}

func bar(i int) error {
	kind := "New"
	if i >= 2 {
		// Create new error with stack trace.
		return tracerr.Errorf(kind, "i = %d", i)
	}
	return bar(i + 1)
}
