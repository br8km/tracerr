package example

import (
	"os"

	"github.com/br8km/tracerr"
)

func main() {
	if err := read(); err != nil {
		// Save output to variable.
		text := tracerr.SprintSource(err)
		os.WriteFile("/tmp/tracerr.log", []byte(text), 0644)
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
