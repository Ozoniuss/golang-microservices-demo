package ports

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Ozoniuss/golang-microservices-demo/file-streamer/internal/decoder"
	"github.com/Ozoniuss/golang-microservices-demo/file-streamer/internal/files"
	public "github.com/Ozoniuss/golang-microservices-demo/file-streamer/pkg"
	"github.com/Ozoniuss/golang-microservices-demo/protobuf/ports/api"
	proto "github.com/Ozoniuss/golang-microservices-demo/protobuf/ports/model"
	log "github.com/Ozoniuss/stdlog"

	"github.com/gin-gonic/gin"
)

func (h *handler) handleLocalPortStream(ctx *gin.Context) {

	filename := ctx.Param("filename")
	if filename == "" {
		err := public.Error{
			Status: http.StatusBadRequest,
			Title:  "File streaming failed",
			Detail: "Missing filename query parameter.",
		}
		ctx.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	f, err := files.OpenFile(filename, h.config.Files)
	if err != nil {
		log.Errf("could not open file with name %s: %w", filename, err)
		clientErr := public.Error{
			Status: http.StatusBadRequest,
			Title:  "File streaming failed",
			Detail: fmt.Sprintf("Could not open file with name %s. Make sure the provided name is valid.", filename),
		}
		ctx.JSON(500, gin.H{
			"error": clientErr,
		})
	}

	buf := make([]byte, h.config.PortService.Decoder.Bufsize)
	var bytesread int
	var readerr error = nil

	dec := decoder.NewPortDecoder()

	count := 0
	client, err := h.client.StreamPorts(ctx)
	if err != nil {
		log.Errf("could not create client: %s", err.Error())
	}

	for {
		bytesread, readerr = f.Read(buf)
		if readerr == io.EOF {
			break
		}

		ports, err := dec.Decode(buf[:bytesread])
		if err != nil {
			panic(err)
		}
		if len(ports) > 0 {
			printMaps(ports)
			list := api.PortList{}
			for k, v := range ports {
				v.Id = k
				list.Ports = append(list.Ports, &v)
			}
			count += len(ports)
			client.Send(&api.StreamPortsRequest{
				Message: &api.StreamPortsRequest_Data{
					Data: &list,
				},
			})
		}
	}
	fmt.Printf("total count: %d", count)

	ctx.Status(200)
}

func printMaps(m map[string]proto.Port) {
	out := ""
	for k := range m {
		out += k + " "
	}
	fmt.Println(out)
}
