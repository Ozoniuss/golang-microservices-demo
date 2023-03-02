package api

import (
	pb "github.com/Ozoniuss/golang-microservices-demo/protobuf/ports/api"

	cfg "github.com/Ozoniuss/golang-microservices-demo/ports/internal/config"
	"github.com/Ozoniuss/golang-microservices-demo/ports/internal/store"
	"github.com/Ozoniuss/golang-microservices-demo/ports/internal/store/inmemory"
)

type PortService struct {
	stoarge store.Store
	pb.UnimplementedPortsServer
}

func NewPortsService(config cfg.Config) (*PortService, error) {

	if config.Database.Inmemory {
		storage := inmemory.NewPortsDatabase()
		return &PortService{
			stoarge: storage,
		}, nil
	}

	// db, err := database.Connect(config.Database)
	// if err != nil {
	// 	return nil, fmt.Errorf("could not connect to database: %w", err)
	// }
	// return &PortService{
	// 	stoarge: db,
	// }, nil
	return nil, nil
}
