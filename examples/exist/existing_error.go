package example

import (
	"os"

	"github.com/br8km/tracerr"
)

func main() {
	if err := read(); err != nil {
		tracerr.PrintSourceColor(err)
	}
}

func read() error {
	return readNonExistent()
}

func readNonExistent() error {
	kind := "NonExist"
	_, err := os.ReadFile("/tmp/non_existent_file")
	// Add stack trace to existing error, no matter if it's nil.
	return tracerr.Wrap(kind, err)
}
