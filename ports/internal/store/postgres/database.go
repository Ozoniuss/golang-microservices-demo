package postgres

import (
	"fmt"

	"github.com/Ozoniuss/golang-microservices-demo/ports/internal/config"
	"github.com/Ozoniuss/golang-microservices-demo/ports/internal/store"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

/*
	See https://gorm.io/docs/query.html#Retrieving-a-single-object for query
	documentation.
*/

func (p *PortsDatabase) Get(id string) (store.Port, error) {
	port := store.Port{}
	err := p.db.First(&port).Error
	if err != nil {
		return port, fmt.Errorf("could not retrieve port: %w", err)
	}
	return port, nil
}
func (p *PortsDatabase) Create(port store.Port) (store.Port, error) {
	// unimplemented
	return store.Port{}, nil
}
func (p *PortsDatabase) Save(port store.Port) (store.Port, error) {
	// unimplemented
	return store.Port{}, nil
}

/*
	See https://gorm.io/docs/create.html#Batch-Insert for batch insert
	documentation.
*/

func (p *PortsDatabase) BatchCreate(ports []store.Port) (int, error) {

	// Update every column on conflict.
	// TODO: return the number of items that did not exist before.
	err := p.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		//DoUpdates: clause.AssignmentColumns([]string{"*"}),
		UpdateAll: true, // Does this update everything?
	}).CreateInBatches(ports, p.config.BatchSize).Error
	if err != nil {
		return 0, fmt.Errorf("could not insert ports: %w", err)
	}

	return len(ports), nil
}
func (p *PortsDatabase) SafeBatchCreate(ports []store.Port) error {

	// Fail if any of the records already exists in the database.
	err := p.db.CreateInBatches(ports, p.config.BatchSize).Error
	if err != nil {
		return fmt.Errorf("could not insert ports: %w", err)
	}

	return nil
}
