package config

type Config struct {
	Server      Server
	PortService PortService
	Files       Files
}

type Server struct {
	Address string
	Port    int32
}

type PortService struct {
	Address string
	Port    int32
	Decoder PortDecoderOptions
}

type PortDecoderOptions struct {
	Bufsize int
}

type Files struct {
	Mount string
}

func newConfig() Config {
	return Config{}
}
