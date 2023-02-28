package postgres

import (
	"fmt"

	cfg "github.com/Ozoniuss/golang-microservices-demo/ports/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// getDSN returns the dsn of the postgres database, as configured by the
// service.
func getDSN(config cfg.Database) string {
	return fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v",
		config.Host, config.Port, config.User, config.Name, config.Password)
}

// Connect establishes a connection to the postgres database.
func Connect(config cfg.Database) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(getDSN(config)))
}
