package influxdb

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/jrlmx2/stockAnalysis/utils/config"
)

//Pool is a wrapper for a golang database/sql.DB object
type Pool struct {
	db *sql.DB
}

var conn client.Client
var bpConf client.BatchPointsConfig

var bp client.BatchPoints

func Setup(conf config.Database) (client.Client, error) {
	// Make client
	var err error

	conn, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     conf.Host,
		Username: conf.User,
		Password: conf.Password,
	})

	if err != nil {
		log.Fatalf("Error: %s", err)
		return nil, err
	}

	bpConf = client.BatchPointsConfig{
		Database:  conf.Schema,
		Precision: "s",
	}

	storeInterval, _ := time.ParseDuration("1s")
	go storePoints(storeInterval)

	return conn, nil
}

func storePoints(storeInterval time.Duration) {

	for {
		var err error

		if bp != nil && len(bp.Points()) > 0 {
			fmt.Println("storing")
			conn.Write(bp)
		}

		bp, err = client.NewBatchPoints(bpConf)
		if err != nil {
			log.Fatal("Error during storepoints function at creating new batchpoints", err)
			return
		}
		time.Sleep(storeInterval)
	}
}

func AddPoint(measurement string, tags map[string]string, fields map[string]interface{}) error {
	pt, err := client.NewPoint(measurement, tags, fields, time.Now())
	fmt.Printf("Added new point %+v\n", pt)
	if err != nil {
		log.Fatal("Error adding datapoint: ", err)
		return err
	}

	bp.AddPoint(pt)
	//conn.Write(bp)
	fmt.Println("point Added")
	return nil
}
