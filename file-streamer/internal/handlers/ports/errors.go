package ports

import (
	public "github.com/Ozoniuss/golang-microservices-demo/file-streamer/pkg"
)

// NewLocalFileStreamingFailedError creates an error that occured when trying
// to stream the content of a local file.
func NewLocalFileStreamingFailedError(status int, detail string) public.Error {
	return public.Error{
		Title:  "Local File Streaming Failed",
		Status: status,
		Detail: detail,
	}
}

// NewLocalFileDebugError creates an error that occured during debugging a
// local ports file.
func NewLocalFileDebugError(status int, detail string) public.Error {
	return public.Error{
		Title:  "Debug File Failed",
		Status: status,
		Detail: detail,
	}
}
