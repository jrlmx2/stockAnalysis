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
	fmt.Printf("Attempting to Query: %s\n\n", query)
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (p *Pool) QueryRow(query string) *sql.Row {
	fmt.Printf("Attempting to QueryRow: %s\n\n", query)
	row := p.db.QueryRow(query)
	fmt.Printf("QueryRow Returned: %+v", row)
	return row
}

func (p *Pool) Exec(query string) (*sql.Result, int64, error) {
	fmt.Printf("Attempting to Exec: %s\n\n", query)
	result, err := p.db.Exec(query)
	if err != nil {
		return nil, 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, 0, nil
	}
	return &result, id, nil
}
