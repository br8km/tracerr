package example

import (
	"fmt"

	"github.com/br8km/tracerr"
)

func main() {
	if err := nilError(); err != nil {
		tracerr.PrintSourceColor(err)
	} else {
		fmt.Println("no error")
	}
}

func nilError() error {
	kind := ""
	return tracerr.Wrap(kind, nil)
}
