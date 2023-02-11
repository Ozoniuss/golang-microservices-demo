package config

import (
	"fmt"
	"os"
	"strings"

	log "golang-microservices-demo/ports/internal/logger"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// defineFlags goes through all the configuration options of the project, and
// defines flags for them that can be supplied by the user.
func defineFlags(configOptions []ConfigOption) error {

	pflag.Bool("write-config", false, "If supplied, the project configuration is written either at the default path, or at the specified location if used in combination with the --config option.")
	pflag.StringP("config", "c", "", "Specifies the exact location of a configuration file. By default, the service attempts to read from ./configs/config.yml")

	for _, opt := range configOptions {
		switch opt.Value.(type) {
		case bool:
			pflag.BoolP(opt.FlagName, opt.Shorthand, opt.Value.(bool), opt.Usage)
		case string:
			pflag.StringP(opt.FlagName, opt.Shorthand, opt.Value.(string), opt.Usage)
		case int:
			pflag.IntP(opt.FlagName, opt.Shorthand, opt.Value.(int), opt.Usage)
		case int32:
			pflag.Int32P(opt.FlagName, opt.Shorthand, opt.Value.(int32), opt.Usage)
		default:
			return fmt.Errorf("Invalid flag value provided for option %s", opt.FlagName)
		}
	}
	return nil
}

// readCustomFlags reads the two custom flags for the config location and
// config overwrite.
func readCustomFlags() (explicitConfigPath string, writeConfig bool, err error) {
	explicitConfigPath, err = pflag.CommandLine.GetString("config")
	if err != nil {
		return
	}
	writeConfig, err = pflag.CommandLine.GetBool("write-config")
	if err != nil {
		return
	}
	return
}

func ParseConfig() (Config, error) {

	appOptions := allOptions()
	c := Config{}

	err := defineFlags(appOptions)
	if err != nil {
		return c, fmt.Errorf("Unable to define flags: %w", err)
	}
	pflag.Parse()

	v := viper.New()

	// Set up the default configuration file options.
	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath("configs")

	// Viper will look for the configuration options specified via environment
	// variables in the form of PORTS_*. The call to AutomaticEnv() makes it
	// no longer necessary to bind to environment variables manually.
	v.SetEnvPrefix("ports")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	for _, opt := range appOptions {

		// "Special" configuration options that can only be set through flags.
		if opt.ConfigKey == "" {
			continue
		}
		v.SetDefault(opt.ConfigKey, opt.Value)

		// Bind to the defined flags. Flags may be left empty.
		if f := pflag.Lookup(opt.FlagName); f != nil {
			err := v.BindPFlag(opt.ConfigKey, f)
			if err != nil {
				return c, fmt.Errorf("could not bind to flag %s: %w", opt.FlagName, err)
			}
		}
	}

	// BindPFlags cannot be used since it uses each flag's full name as the
	// config key (see function documentation).
	//v.BindPFlags(pflag.CommandLine)

	explicitConfigPath, writeConfig, err := readCustomFlags()
	if err != nil {
		return c, fmt.Errorf("Custom flags were not parsed correctly: %w", err)
	}

	if explicitConfigPath != "" {
		v.SetConfigFile(explicitConfigPath)
	}

	if err := v.ReadInConfig(); err != nil {
		// For some reason errors.Is() doesn't work with this viper error.
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warnln("Config file not found. Use --write-config to create it")

			// And for whatever dumb reason viper doesn't throw the same error
			// if the config file was explicitly set.
		} else if _, ok := err.(*os.PathError); ok {
			log.Warnln("Config file not found. Use --write-config to create it")
		} else {

		}
	}

	if writeConfig != false {
		log.Infoln("Writing configuration file.")
		err := v.WriteConfig()

		// For some reason WriteConfig() doesn't work with the default location.
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			v.SafeWriteConfig()
		}
	}

	v.Unmarshal(&c)
	return c, nil
}
