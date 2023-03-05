package files

import (
	"fmt"
	"os"
	"path"

	"github.com/Ozoniuss/golang-microservices-demo/file-streamer/internal/config"
)

// localSystem provides operation for interacting with the server's local
// system.
type localSystem struct {
	mountPath string
}

func newLocalSystem(config config.Files) *localSystem {
	return &localSystem{
		mountPath: config.Mount,
	}
}

func (l *localSystem) OpenFile(filename string) (*os.File, error) {

	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("could not open working directory: %w", err)
	}

	// Generate the full local file path, given the mount point of the local
	// files.
	filepath := path.Join(wd, l.mountPath, filename)
	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	return f, nil
}

func (l *localSystem) ListFileNames(dir string) ([]string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("could not open working directory: %w", err)
	}

	filepath := path.Join(wd, l.mountPath, dir)
	stat, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("the provided path \"%s\" does not exist on the server's filesystem.", filepath)
		} else {
			return nil, fmt.Errorf("error getting stats for %s: %w", filepath, err)
		}
	}
	if !stat.IsDir() {
		return nil, fmt.Errorf("the provided path \"%s\" is not a directory on the server's filesystem.", filepath)
	}

	entries, err := os.ReadDir(filepath)
	if err != nil {
		return nil, fmt.Errorf("could not read files at path \"%s\": %w", filepath, err)
	}

	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			names = append(names, entry.Name())
		}
	}
	return names, nil
}
