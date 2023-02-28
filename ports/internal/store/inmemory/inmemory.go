package inmemory

import (
	"fmt"
	"sync"

	"github.com/Ozoniuss/golang-microservices-demo/ports/internal/store"
)

type PortsDatabase struct {
	// use a map to get in O(1)
	records map[string]store.Port
	// sync concurrent access to the database accross multiple threads
	mutex *sync.RWMutex
}

func NewPortsDatabase() store.Store {
	return &PortsDatabase{
		records: make(map[string]store.Port),
	}
}

func (pd *PortsDatabase) Get(id string) (store.Port, error) {
	pd.mutex.RLock()
	defer pd.mutex.RUnlock()
	port, ok := pd.records[id]
	if !ok {
		return store.Port{}, fmt.Errorf("port with id %s not found.", id)
	}
	return port, nil
}

func (pd *PortsDatabase) Create(port store.Port) (store.Port, error) {
	pd.mutex.Lock()
	defer pd.mutex.Unlock()
	_, found := pd.records[port.Id]
	if found {
		return store.Port{}, fmt.Errorf("port with id %s already exists.", port.Id)
	}
	pd.records[port.Id] = port
	return port, nil
}

func (pd *PortsDatabase) Save(port store.Port) (store.Port, error) {
	pd.mutex.Lock()
	defer pd.mutex.Unlock()
	old, found := pd.records[port.Id]
	pd.records[port.Id] = port
	if found {
		return old, nil
	}
	return port, nil
}

func (pd *PortsDatabase) BatchCreate(ports []store.Port) (int, error) {
	// unimplemented
	return 0, nil
}

func (pd *PortsDatabase) SafeBatchCreate(port []store.Port) error {
	// unimplemented
	return nil
}

func (pd *PortsDatabase) toString() string {
	return "ports"
}
