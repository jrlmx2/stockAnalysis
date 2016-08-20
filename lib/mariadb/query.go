package mariadb

import (
	"database/sql"
	"fmt"
)

type queryable interface {
	Table() string
	Data() string
	NewBunch(*sql.Rows) []*queryable
	NewOne(*sql.Row) *queryable
}

const (
	query  = "SELECT %s FROM %s%s"
	insert = "INSERT INTO %s VALUES %s"
)

func (p *Pool) Select(q queryable, where, cols string) ([]*queryable, error) {
	rows, err := p.db.Query(fmt.Sprintf(query, cols, q.Table(), where))
	if err != nil {
		return nil, err
	}

	return q.NewBunch(rows), nil
}

func (p *Pool) SelectOne(q queryable, where, cols string) (*queryable, error) {
	return q.NewOne(p.db.QueryRow(fmt.Sprintf(query, cols, q.Table(), where))), nil
}

func (p *Pool) SaveOne(q queryable) (*queryable, int64, error) {
	result, err := p.db.Exec(fmt.Sprintf(insert, q.Table(), q.Data()))
	if err != nil {
		return nil, 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, 0, nil
	}
	return &q, id, nil
}

func (p *Pool) SaveAll(q []queryable) ([]queryable, []int64, error) {
	var result sql.Result
	var err error
	ids := make([]int64, 0)
	var id int64
	for _, r := range q {
		result, err = p.db.Exec(fmt.Sprintf(insert, r.Table(), r.Data()))
		if err != nil {
			return nil, nil, err
		}
		id, err = result.LastInsertId()
		ids = append(ids, id)
	}

	return q, ids, nil
}
