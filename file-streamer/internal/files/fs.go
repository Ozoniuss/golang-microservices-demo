package files

import "os"

// FileSystem defines the possible interactions with a file system, regardless
// of the actual technology that powers the file system.
type FileSystem interface {
	// OpenFile opens the file and returns an *os.File object, which implements
	// multiple interfaces such as io.Reader, io.Writer and io.Seeker.
	OpenFile(filename string) (*os.File, error)
	// ListFileNames lists the names of the files at the given path.
	ListFileNames(dir string) ([]string, error)
}
