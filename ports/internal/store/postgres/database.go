package postgres

import (
	"fmt"

	"github.com/Ozoniuss/golang-microservices-demo/ports/internal/config"
	"github.com/Ozoniuss/golang-microservices-demo/ports/internal/store"
	"gorm.io/gorm"
)

// PortsDatabase holds the connection to the postgres database server, as well
// as the database configuration.
type PortsDatabase struct {
	db     *gorm.DB
	config config.Database
}

func NewPortsDatabase(config config.Database) (store.Store, error) {
	db, err := connect(config)
	if err != nil {
		return nil, fmt.Errorf("could not connect to postgres database: %w", err)
	}
	p := PortsDatabase{
		db:     db,
		config: config,
	}
	return &p, nil
}

// unimplemented

func (p *PortsDatabase) Get(id string) (store.Port, error) {
	return store.Port{}, nil
}
func (p *PortsDatabase) Create(port store.Port) (store.Port, error) {
	return store.Port{}, nil
}
func (p *PortsDatabase) Save(port store.Port) (store.Port, error) {
	return store.Port{}, nil
}
func (p *PortsDatabase) BatchCreate(ports []store.Port) (int, error) {
	return 0, nil
}
func (p *PortsDatabase) SafeBatchCreate(ports []store.Port) error {
	return nil
}
