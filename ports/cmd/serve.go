package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/Ozoniuss/golang-microservices-demo/ports/internal/config"
	log "github.com/Ozoniuss/stdlog"

	"google.golang.org/grpc"
)

func serveGRPC(grpcServer *grpc.Server, config config.Server) error {

	l, err := net.Listen(config.Network, fmt.Sprintf(":%d", config.Port))
	if err != nil {
		return fmt.Errorf("could not start ports server: %w", err)
	}
	defer l.Close()
	log.Infoln("Ports server running.")

	errc := make(chan error)
	go func() {
		err := grpcServer.Serve(l)
		if err != nil {
			errc <- err
		}
	}()

	// Capture termination sigals to allow for graceful shutdown.
	sigc := make(chan os.Signal, 1)

	signal.Notify(sigc, os.Interrupt)
	signal.Notify(sigc, os.Kill)

	select {

	// Captures any errors received in the grpc.Serve() function above.
	case err := <-errc:
		return err

	// Captures termination signals.
	case <-sigc:

		// Used to mark the termination of graceful shutdown.
		gracefulc := make(chan struct{})
		go func() {
			log.Infoln("Shutting down gracefully, stop again to force...")
			grpcServer.GracefulStop()
			// Signal the select statement to no longer listen on the force
			// shutdown channel.
			close(gracefulc)
		}()

		select {

		// Graceful shutdown complete.
		case <-gracefulc:
			log.Infoln("Server shut down gracefully.")

		// Forced shutdown.
		case <-sigc:
			grpcServer.Stop()
			log.Infoln("Server shut down.")

		// Timeout, do not allow the server to shut down forever.
		case <-time.After(config.ShutdownTimeout):
			grpcServer.Stop()
			log.Infoln("Server shut down after timeout.")
		}
		return nil
	}
}
