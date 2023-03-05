package ports

import (
	pb "github.com/Ozoniuss/golang-microservices-demo/protobuf/ports/api"
	proto "github.com/Ozoniuss/golang-microservices-demo/protobuf/ports/model"

	"github.com/Ozoniuss/golang-microservices-demo/file-streamer/internal/config"
	"github.com/Ozoniuss/golang-microservices-demo/file-streamer/internal/decoder"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type handler struct {
	client  pb.PortsClient
	decoder decoder.MapDecoder[string, proto.Port]
	config  config.Config
}

func NewHandler(conn *grpc.ClientConn, config config.Config) *handler {
	return &handler{
		client:  pb.NewPortsClient(conn),
		decoder: decoder.NewPortDecoder(),
		config:  config,
	}
}

func RegisterPorts(router *gin.Engine, config config.Config, conn *grpc.ClientConn) {

	h := NewHandler(conn, config)

	filestreamRouter := router.Group("/api/ports/local")
	{
		filestreamRouter.POST("/::filename", h.handleLocalPortStream)
	}
	debugRouter := router.Group("/api/ports/debug")
	{
		debugRouter.GET("/:filename", h.handleDebugLocalFile)
	}
}
