package mariadb

import (
	"database/sql"
	"fmt"

	"github.com/jrlmx2/stockAnalysis/lib/config"
	_ "github.com/ziutek/mymysql/godrv"
)

const connectionString = "tcp:%s*%s/%s/%s"

type Pool struct {
	db *sql.DB
}

func NewPool(conf config.Database) (*Pool, error) {

	db, err := sql.Open("mymysql", fmt.Sprintf(connectionString, conf.Host, conf.Schema, conf.User, conf.Password))
	if err != nil {
		fmt.Printf("\n\nConnection error: %s", err)
		return nil, err
	}
	defer db.Close()

	db.SetMaxOpenConns(10)

	return &Pool{db: db}, nil
}
