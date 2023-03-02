package api

import (
	"fmt"
	"io"

	pb "github.com/Ozoniuss/golang-microservices-demo/protobuf/ports/api"

	log "github.com/Ozoniuss/stdlog"
)

func (serv *PortService) StreamPorts(src pb.Ports_StreamPortsServer) error {
	for {
		rr, err := src.Recv()
		if err == io.EOF {
			log.Infoln("Client has closed connection.")
			return nil
		}
		fmt.Println(rr)
	}
}
