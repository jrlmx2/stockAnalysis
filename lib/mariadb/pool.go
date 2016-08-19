package mariadb

import (
	"fmt"

	"github.com/jrlmx2/stockAnalysis/lib/config"
	"github.com/ziutek/mymysql/mysql"
)

type Pool struct {
}

func NewPool(conf config.Database) (*Pool, error) {
	db := mysql.New("tcp", "", conf.Host, conf.User, conf.Password, conf.Schema)
	if err := db.Connect(); err != nil {
		fmt.Println(err)
	}

	var query = "INSERT INTO symbols VALUES ('TVIX');"

	rows, results, err := db.Query(query, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Rows: %+v\nResults: %+v\n", rows, results)

	return &Pool{}, nil
}
