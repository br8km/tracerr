package example

import (
	"fmt"
	"os"

	"github.com/br8km/tracerr"
)

func main() {
	if err := read(); err != nil {
		// Dump raw stack trace.
		frames := tracerr.StackTrace(err)
		fmt.Printf("%#v\n", frames)
	}
}

func read() error {
	return readNonExistent()
}

func readNonExistent() error {
	kind := "NonExist"
	_, err := os.ReadFile("/tmp/non_existent_file")
	return tracerr.Wrap(kind, err)
}
