package config

import "time"

const (
	KiB = 1024
	MiB = 1024 * KiB
	GiB = 1024 * MiB
	TiB = 1024 * GiB
	PiB = 1024 * TiB

	ConnectionType = "tcp"
	Port           = "9092"

	OperationsTimeout = 30 * time.Second //max time to allow service methods to work after closing server

	GRPCMaxMessageSize = 50 * MiB

	JSONIndent  = 4
	JSONNewLine = 2

	InMemoryFlag = true //true for inmemory database, false for postgers database
)
