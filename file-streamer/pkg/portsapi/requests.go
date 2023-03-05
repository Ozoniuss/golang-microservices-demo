package portsapi

// DebugRequest holds the information that can be sent to a debug request.
type DebugRequest struct {
	BufferSize *int `json:"buffer_size,omitempty"`
}
