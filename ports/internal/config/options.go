package config

type ConfigOption struct {
	FlagName  string
	Shorthand string
	Value     any
	Usage     string
	ConfigKey string
}

func databaseOptions() []ConfigOption {
	return []ConfigOption{
		{FlagName: "db-name", Shorthand: "", Value: "ports_db", ConfigKey: "database.name",
			Usage: "Specifies the name of the ports database"},
		{FlagName: "db-host", Shorthand: "", Value: "127.0.0.1", ConfigKey: "database.host",
			Usage: "Specifies the address on which the ports database listens for connections"},
		{FlagName: "db-port", Shorthand: "", Value: int32(5432), ConfigKey: "database.port",
			Usage: "Specifies the port on which the ports database listens for connections"},
		{FlagName: "db-user", Shorthand: "", Value: "ports_user", ConfigKey: "database.user",
			Usage: "Specifies the user which connects to the ports database"},
		{FlagName: "db-password", Shorthand: "", Value: "userpassword", ConfigKey: "database.password",
			Usage: "Specifies the password of the user which connects to the ports database"},
	}
}

func serverOptions() []ConfigOption {
	return []ConfigOption{
		{FlagName: "server-address", Shorthand: "", Value: "127.0.0.1", ConfigKey: "server.address",
			Usage: "Specifies the address on which the ports service listens for incoming calls"},
		{FlagName: "server-port", Shorthand: "", Value: int32(9000), ConfigKey: "server.port",
			Usage: "Specifies the port on which the ports service listens for incoming calls"},
	}
}

func allOptions() []ConfigOption {
	opts := make([]ConfigOption, 0)
	opts = append(opts, databaseOptions()...)
	opts = append(opts, serverOptions()...)
	return opts
}
