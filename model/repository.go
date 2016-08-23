package model

import "github.com/jrlmx2/stockAnalysis/utils/mariadb"

var repository *mariadb.Pool

func SetRepository(conn *mariadb.Pool) {
	repository = conn
}
