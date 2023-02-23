package config

import (
	"fmt"

	"github.com/Ozoniuss/goutils/memsize"

	cfg "github.com/Ozoniuss/configer"
)

type ConfigOption struct {
	FlagName  string
	Shorthand string
	Value     any
	Usage     string
	ConfigKey string
}

func serverOptions() []cfg.ConfigOption {
	return []cfg.ConfigOption{
		{FlagName: "server-address", Shorthand: "", Value: "127.0.0.1", ConfigKey: "server.address",
			Usage: "Specifies the address on which the file streamer api listens for incoming requests"},
		{FlagName: "server-port", Shorthand: "", Value: int32(7070), ConfigKey: "server.port",
			Usage: "Specifies the port on which the file streamer api listens for incoming requests"},
	}
}

func portServiceOptions() []cfg.ConfigOption {
	return []cfg.ConfigOption{
		{FlagName: "ports-address", Shorthand: "", Value: "localhost", ConfigKey: "portservice.address",
			Usage: "Specifies the address on which the ports service listens for incoming calls"},
		{FlagName: "ports-port", Shorthand: "", Value: int32(9000), ConfigKey: "portservice.port",
			Usage: "Specifies the port on which the ports service listens for incoming calls"},
		{FlagName: "port-bufsize", Shorthand: "", Value: int(32 * memsize.KiB), ConfigKey: "portservice.decoder.bufsize",
			Usage: "Specifies the buffer size of the ports decoder"},
	}
}

func filesOptions() []cfg.ConfigOption {
	return []cfg.ConfigOption{
		{FlagName: "ports-file-mount", Shorthand: "", Value: "./assets", ConfigKey: "files.mount",
			Usage: "Specifies the mount point of the files that contain data for local algorithms."},
	}
}

func allOptions() []cfg.ConfigOption {
	opts := make([]cfg.ConfigOption, 0)
	opts = append(opts, serverOptions()...)
	opts = append(opts, portServiceOptions()...)
	opts = append(opts, filesOptions()...)
	return opts
}

func ParseConfig() (Config, error) {
	c := newConfig()

	parserOptions := []cfg.ParserOption{
		cfg.WithConfigName("config"),
		cfg.WithConfigType("yml"),
		cfg.WithConfigPath("./configs"),
		cfg.WithEnvPrefix("PORTS"),
		cfg.WithEnvKeyReplacer("_"),
		cfg.WithWriteFlag(),
	}

	err := cfg.NewConfig(&c, allOptions(), parserOptions...)
	if err != nil {
		return newConfig(), fmt.Errorf("could not create config: %w", err)
	}
	return c, nil
}
