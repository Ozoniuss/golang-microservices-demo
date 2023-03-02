package main

import (
	"fmt"
	"os"

	cfg "github.com/Ozoniuss/golang-microservices-demo/file-streamer/internal/config"
	"github.com/Ozoniuss/golang-microservices-demo/file-streamer/internal/server"

	log "github.com/Ozoniuss/stdlog"
)

func run() error {
	config, err := cfg.ParseConfig()
	if err != nil {
		return fmt.Errorf("could not parse config: %w", err)
	}

	engine, err := server.NewEngine(config)
	engine.Run(fmt.Sprintf("%s:%d", config.Server.Address, config.Server.Port))

	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Errf("Error running api: %s", err.Error())
		os.Exit(1)
	}
}
