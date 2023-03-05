package portsapi

// DebugRequest holds the information that can be sent to a debug request.
type DebugRequest struct {
	BufferSize *int `json:"buffer_size,omitempty"`
}

// DebugResponse holds the information that can be received as a response to
// a debug request.
type DebugResponse struct {
	// The number of ports read from the file.
	PortsRead int `json:"ports_read"`
	// The number of reads through the ports file.
	BufferPasses int `json:"buffer_passes"`
}
