package api

import (
	"fmt"

	pb "github.com/Ozoniuss/golang-microservices-demo/protobuf/ports/api"

	cfg "github.com/Ozoniuss/golang-microservices-demo/ports/internal/config"
	"github.com/Ozoniuss/golang-microservices-demo/ports/internal/store"
	"github.com/Ozoniuss/golang-microservices-demo/ports/internal/store/inmemory"
	"github.com/Ozoniuss/golang-microservices-demo/ports/internal/store/postgres"
)

type PortService struct {
	stoarge store.Store
	pb.UnimplementedPortsServer
}

func NewPortsService(config cfg.Config) (*PortService, error) {

	var storage store.Store
	var err error

	if config.Database.Inmemory {
		storage = inmemory.NewPortsDatabase()
		return &PortService{
			stoarge: storage,
		}, nil
	} else {
		storage, err = postgres.NewPortsDatabase(config.Database)
		if err != nil {
			return nil, fmt.Errorf("could not initialize ports postgres database: %w", err)
		}
	}

	return &PortService{
		stoarge: storage,
	}, nil
}
