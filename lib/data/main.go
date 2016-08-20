package main

import (
	"fmt"
	"time"

	"github.com/jrlmx2/stockAnalysis/lib/config"
	"github.com/jrlmx2/stockAnalysis/lib/data/streaming"
	"github.com/jrlmx2/stockAnalysis/lib/logger"
	"github.com/jrlmx2/stockAnalysis/lib/mariadb"
	"github.com/jrlmx2/stockAnalysis/lib/oauth"
)

func main() {
	conf := config.ReadConfig()
	fmt.Printf("\n\n%+v\n\n", conf)

	logger, _ := log.NewLogger(conf.Logger.Name, conf.Logger.Format, conf.Logger.File, conf.Logger.Level)
	fmt.Printf("%+v", logger)

	oauthWrapper.SetCredentials(conf.Server["tradeking"].OAuthToken, conf.Server["tradeking"].OAuthSecret)
	oauthWrapper.SetClient(conf.Server["tradeking"].Key, conf.Server["tradeking"].Secret)

	pool, err := mariadb.NewPool(conf.Database)
	if err != nil {
		panic(err)
	}

	pool.SaveOne(*streaming.TradeDetails{Last: 34.2, Symbol: "TVIX", SymbolID: 5, Timestamp: time.Now().Unix(), Amount: 5000, Vwap: 33.2})
	//streaming.OpenStream([]string{"symbols=TVIX"})
}
