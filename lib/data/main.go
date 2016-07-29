package main

import (
  "fmt"
  "github.com/jrlmx2/stockAnalysis/lib/data/config"
  "github.com/jrlmx2/stockAnalysis/lib/logger"
)

func main() {
  conf := config.ReadConfig()
  fmt.Printf("%+v", conf)

  logger, _ := log.NewLogger(conf.Logger.Name, conf.Logger.Format, conf.Logger.File, conf.Logger.Level)
  fmt.Printf("%+v",logger)
}
