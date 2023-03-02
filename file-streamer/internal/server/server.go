package server

import (
	"fmt"

	cfg "github.com/Ozoniuss/golang-microservices-demo/file-streamer/internal/config"
	porthandler "github.com/Ozoniuss/golang-microservices-demo/file-streamer/internal/handlers/ports"

	log "github.com/Ozoniuss/stdlog"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewEngine(config cfg.Config) (*gin.Engine, error) {

	r := gin.Default()

	{
		log.Infoln("Dialing ports service...")
		conn, err := dialGRPC(config.PortService.Address, config.PortService.Port)
		if err != nil {
			return nil, fmt.Errorf("could not establish connection with the ports service: %w", err)
		}
		log.Infoln("Ports client created.")
		porthandler.RegisterPorts(r, config, conn)
		log.Infoln("Ports handler registered.")
	}

	return r, nil
}

func dialGRPC(addr string, port int32) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption

	// Secure connection not implemented yet
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	return grpc.Dial(fmt.Sprintf("%s:%d", addr, port), opts...)
}
