package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// Config type wraps all the potential fields a configuration file might have
type Config struct {
	API      map[string]API
	Logger   LogConfig
	Database Database
	Server   Server
	Stocks   Stocks
}

// LogConfig Describes the options used to setup the logger
type LogConfig struct {
	Name   string
	Level  string
	File   string
	Format string
}

// API contains options and authentication values for API connectivity and query
type API struct {
	Key         string
	Secret      string
	OAuthToken  string
	OAuthSecret string
}

// Database contains options for database authentication
type Database struct {
	User     string
	Password string
	Host     string
	Schema   string
}

type Stocks struct {
	Symbols string
}

// Server Describies
type Server struct {
	Address string
}

// ReadConfigPath reads a file into Config struct
func ReadConfigPath(file string) *Config {
	fmt.Printf("\n Reading: %s into %+v", file, Database{})
	if _, err := os.Stat(file); err != nil {
		fmt.Printf("\n Error reading config file %+v\n", err)
		panic(err)
	}

	var config Config
	if _, err := toml.DecodeFile(file, &config); err != nil {
		fmt.Printf("\n Error decoding config file %+v\n", err)
		panic(err)
	}
	fmt.Printf("\n Read: %s into %+v", file, config)
	return &config
}

// ReadConfig reads the command line -c filepath into a Config struct
func ReadConfig() (*Config, string) {
	configFile := flag.String("c", "", "Configuration file")
	fmt.Printf("%+v", *configFile)
	if !flag.Parsed() {
		flag.Parse()
	}

	if _, err := os.Stat(*configFile); err != nil {
		fmt.Printf("%+v\n", err)
		//handle file doesn't exist error
	}

	var config Config
	if _, err := toml.DecodeFile(*configFile, &config); err != nil {
		fmt.Printf("%+v\n", err)
		//handle config parsing error
	}

	return &config, *configFile
}
