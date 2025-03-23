// This file is added for purpose of having an example of different path in tests.
package tracerr_test

import (
	"github.com/br8km/tracerr"
)

var KIND string = "kind"

func addFrameA(kind string, message string) error {
	return addFrameB(kind, message)
}

func addFrameB(kind string, message string) error {
	return addFrameC(kind, message)
}

func addFrameC(kind string, message string) error {
	return tracerr.New(kind, message)
}
