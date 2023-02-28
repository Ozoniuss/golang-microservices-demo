package main

import (
	"fmt"

	"github.com/Ozoniuss/golang-microservices-demo/ports/internal/api"
	"github.com/Ozoniuss/golang-microservices-demo/ports/internal/config"

	log "github.com/Ozoniuss/stdlog"

	"os"

	pb "github.com/Ozoniuss/golang-microservices-demo/protobuf/ports/api"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// run starts the server and executes the main code.
func run() error {

	config, err := config.ParseConfig()
	if err != nil {
		return err
	}
	log.Infof("Parsed configuration options: %+v", config)

	server, err := api.NewPortsService(config)
	if err != nil {
		return fmt.Errorf("could not create ports service: %w", err)
	}
	log.Infoln("Ports service created.")

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_recovery.UnaryServerInterceptor(),
			),
		),
	)
	pb.RegisterPortsServer(grpcServer, server)
	log.Infoln("Ports service registered to the grpc server.")

	// For debugging.
	reflection.Register(grpcServer)
	log.Infoln("Reflection registered.")

	err = serveGRPC(grpcServer, config.Server)
	if err != nil {
		return fmt.Errorf("serving grpc: %w", err)
	}

	return nil
}

func main() {

	// os.Exit() doesn't allow deferred functions to execute once called. This
	// pattern allows executing deferred functions.
	//
	// To provide custom exit codes, the os.Exit() call can be deferred instead,
	// and the code set in the main function when an error is captured. In this
	// case though this is good enough.
	if err := run(); err != nil {
		log.Errf("Error running ports service: %s", err.Error())
		os.Exit(1)
	}
}
