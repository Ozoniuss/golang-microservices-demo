package ports

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Ozoniuss/golang-microservices-demo/file-streamer/internal/decoder"
	"github.com/Ozoniuss/golang-microservices-demo/file-streamer/internal/files"
	"github.com/Ozoniuss/golang-microservices-demo/protobuf/ports/api"
	"github.com/gin-gonic/gin"

	portsapi "github.com/Ozoniuss/golang-microservices-demo/file-streamer/pkg/portsapi"
)

func (h *handler) handleDebugLocalFile(ctx *gin.Context) {

	filename := ctx.Param("filename")
	if filename == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": NewLocalFileDebugError(http.StatusBadRequest, "Missing filename query parameter."),
		})
		return
	}

	// Debug opening file.
	f, err := files.OpenFile(filename, h.config.Files)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": NewLocalFileStreamingFailedError(http.StatusOK,
				fmt.Sprintf("Error opening file %s: %s", filename, err.Error())),
		})
	}

	var req portsapi.DebugRequest
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": NewLocalFileDebugError(http.StatusBadGateway,
				fmt.Sprintf("invalid query parameters: %s", err.Error())),
		})
		return
	}

	bufsize := h.config.PortService.Decoder.Bufsize
	if req.BufferSize != nil {
		bufsize = *req.BufferSize
	}

	buf := make([]byte, bufsize)
	var bytesread int
	var readerr error = nil

	dec := decoder.NewPortDecoder()

	count := 0
	passes := 0

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
			list := api.PortList{}
			for k, v := range ports {
				v.Id = k
				list.Ports = append(list.Ports, &v)
			}
			count += len(ports)
		}
		passes += 1
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Ports read":    count,
		"Buffer passes": passes,
	})
}
