package files

import (
	"fmt"
	"os"
	"path"

	"github.com/Ozoniuss/golang-microservices-demo/file-streamer/internal/config"
)

// OpenFile opens the file with the provided filename, given the files
// configuration.
func OpenFile(filename string, config config.Files) (*os.File, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("could not open working directory: %w", err)
	}
	filepath := path.Join(wd, config.Mount, filename)
	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	return f, nil
}
