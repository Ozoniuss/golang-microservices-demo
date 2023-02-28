package store

// Store provides methods to describe the possible interactions with a Ports
// database. Each storage client must implement this interface.
type Store interface {
	// Retrieve a single port. Returns an error if the port doesn't exist.
	Get(id string) (Port, error)
	// Creates a single port. Returns the port that was created, and an error
	// if a port with the same id already exists.
	Create(port Port) (Port, error)
	// Creates a single port, or updates it if one with the same id already
	// exists. If a port already existed, returns the old ported, otherwise
	// returns the new port.
	Save(port Port) (Port, error)
	// Create multiple ports. If a single port cannot be inserted, skips it.
	// Returns the number of ports that were inserted.
	BatchCreate(ports []Port) (int, error)
	// Create multiple ports. If a single port cannot be inserted, returns an
	// error.
	SafeBatchCreate(ports []Port) error
}
