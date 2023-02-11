## Config

This package handles the configuration of the project. It uses [viper](https://github.com/spf13/viper) as the configuration solution for the project, and the [pflag](https://github.com/spf13/pflag) library to work with flags.

I have divided the configuration options in two different types:

- Standard project options, which viper stores as a key-value pair. Includes options such as server and database connectivity details, API keys, secrets, users, timeouts, buffer sizes etc. These may or may not be associated with a flag, but always have a viper bind. They are usually specified through environment variables, although they can also be specified via a configuration file, or default values. Those are often found in most projects;
- Special project options, which can only be specified via flags. This is the options that I considered relevant for the project setup, but not relevant to the project itself once it is running. This is not something standardized, just an approach that I considered best-suited for options such as the location of the configuration file, and whether or not the configuration file should be overwritten.

The configuration setup here is fairly simple and mostly inspired from the viper documentation. However, viper is fairly configurable and enables writing a configuration template that can be used accorss multiple similar projects. I didn't really want to do that for this demo, as it seemed to much unnecessary work.

I do want to note that there are some functions in the viper librarly whose behaviour I found questionable. If interested, I documented them in some comments of the [parse.go](./parse.go) file. There's probably a reason behind it, but I didn't bother to investigate.