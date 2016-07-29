package config

import(
  "os"
  "fmt"
  "flag"
  "github.com/BurntSushi/toml"
)

type Config struct{
  Server  map[string]Api
  Logger  LogConfig
}

type LogConfig struct {
  Name string
  Level string
  File string
  Format string
}

type Api struct {
  Key string
  Secret string
  OAuthToken string
  OAuthSecret string
}

// Reads info from config file
func ReadConfig() Config {
  configFile := flag.String("c", "", "Configuration file")
  fmt.Printf("%+v",*configFile)
  if !flag.Parsed() {
    flag.Parse()
  }

	if 	_, err := os.Stat(*configFile); err != nil {
    fmt.Printf("%+v\n",err)
    //handle file doesn't exist error
	}

	var config Config
	if _, err := toml.DecodeFile(*configFile, &config); err != nil {
    fmt.Printf("%+v\n",err)
    //handle config parsing error
	}

	return config
}
