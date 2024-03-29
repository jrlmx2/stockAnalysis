package mariadb

import (
	"database/sql"
	"fmt"

	"github.com/jrlmx2/stockAnalysis/utils/config"
	_ "github.com/ziutek/mymysql/godrv"
)

const connectionString = "tcp:%s*%s/%s/%s"

var connectString = ""

//Pool is a wrapper for a golang database/sql.DB object
type Pool struct {
	db *sql.DB
}

func (p *Pool) Close() error {
	return p.db.Close()
}

//NewPool wraps the database connection
func NewPool(conf config.Database) (*Pool, error) {
	if connectString == "" {
		connectString = fmt.Sprintf(connectionString, conf.Host, conf.Schema, conf.User, conf.Password)
		fmt.Printf(connectString)
	}

	db, err := sql.Open("mymysql", connectString)
	if err != nil {
		fmt.Printf("\n\nConnection error: %s", err)
		return nil, err
	}

	db.SetMaxOpenConns(20)

	return &Pool{db: db}, nil
}
