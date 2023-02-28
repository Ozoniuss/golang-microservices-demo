package config

import "time"

type Config struct {
	Server   Server
	Database Database
}

type Database struct {
	Inmemory bool
	Host     string
	Port     int32
	User     string
	Name     string
	Password string
}

type Server struct {
	Address         string
	Port            int32
	Network         string
	ShutdownTimeout time.Duration
}

func newConfig() Config {
	c := Config{}
	return c
}
