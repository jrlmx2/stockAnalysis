package mariadb

import (
	"database/sql"
	"fmt"
)

const (
	query  = "SELECT %s FROM %s%s"
	insert = "INSERT INTO %s VALUES %s"
)

func (p *Pool) Query(query string) (*sql.Rows, error) {
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (p *Pool) QueryRow(query string) *sql.Row {
	return p.db.QueryRow(query)
}

func (p *Pool) Exec(query string) (*sql.Result, int64, error) {
	result, err := p.db.Exec(query)
	fmt.Printf("%s\n", err)
	if err != nil {
		return nil, 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, 0, nil
	}
	return &result, id, nil
}
