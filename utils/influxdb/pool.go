package influxdb

import (
	"database/sql"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/jrlmx2/stockAnalysis/utils/config"
	"github.com/jrlmx2/stockAnalysis/utils/logger"
)

//Pool is a wrapper for a golang database/sql.DB object
type Pool struct {
	db *sql.DB
}

var conn client.Client
var bpConf map[string]chan client.Point
var schema string
var running bool
var log *logger.Logger
var bp BatchPoints

func Setup(conf config.Database, logConf config.LogConfig) (client.Client, error) {
	// Make client
	var err error

	conn, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     conf.Host,
		Username: conf.User,
		Password: conf.Password,
	})

	if err != nil {
		log.Fatal("Error: %s", err)
		return nil, err
	}

	log, err = logger.NewLogger("influx", logConf)
	if err != nil {
		return nil, err
	}

	schema = conf.Schema
	bp = BatchPoints{}
	bp.addConfigType("stocks")

	running = true
	return conn, nil
}

func IsRunning() bool {
	return running
}

// queryDB convenience function to query the database
func QueryDB(cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: "markets",
	}
	if response, err := conn.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}
