package main

import (
	"fmt"

	"github.com/jrlmx2/stockAnalysis/lib/config"
	"github.com/jrlmx2/stockAnalysis/lib/logger"
	"github.com/jrlmx2/stockAnalysis/lib/mariadb"
	"github.com/jrlmx2/stockAnalysis/lib/oauth"
)

func main() {
	conf := config.ReadConfig()
	fmt.Printf("%+v", conf)

	logger, _ := log.NewLogger(conf.Logger.Name, conf.Logger.Format, conf.Logger.File, conf.Logger.Level)
	fmt.Printf("%+v", logger)

	oauthWrapper.SetCredentials(conf.Server["tradeking"].OAuthToken, conf.Server["tradeking"].OAuthSecret)
	oauthWrapper.SetClient(conf.Server["tradeking"].Key, conf.Server["tradeking"].Secret)

	_, _ = mariadb.NewPool(conf.Database)

	//streaming.OpenStream([]string{"symbols=TVIX"})
}
