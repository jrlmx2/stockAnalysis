package model

import (
	"fmt"
	"testing"

	"github.com/jrlmx2/stockAnalysis/utils/config"
	"github.com/jrlmx2/stockAnalysis/utils/influxdb"
	"github.com/jrlmx2/stockAnalysis/utils/mariadb"
)

func setup() {
	conf := config.ReadConfigPath("./test.conf")

	pool, err := mariadb.NewPool(conf.Database)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	SetRepository(pool)
}
func setupInflux() {
	conf := config.ReadConfigPath("./test.conf")
	influxdb.Setup(conf.InfluxDatabase)
}

func TestTradeSave001(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	setup()
	setupInflux()

	fmt.Println("Making trade")
	tr := NewTrade("UVXY")
	tr.Last = 23.4
	tr.TradedVolume = 1000
	tr.VolumeWeightedAverage = 22.2

	fmt.Println("Saving")
	tr.Save()

	fmt.Println("Saved")
}
