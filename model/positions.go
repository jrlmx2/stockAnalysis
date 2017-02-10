package model

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	positionImport = "INSERT INTO positions (id, symbol, shares, cost_to_initiate, strike_price, purchased, exit, active) VALUES (NULL, %d, %d, %f, %f, %s, %s, %b)"
	positionInsert = "INSERT INTO positions (id, symbol, shares, cost_to_initiate, strike_price) VALUES (NULL, %d, %d, %f, %f)"
	positionSold   = "UPDATE positions SET active=0 WHERE id=%d"
	positionSelect = "SELECT * FROM positions WHERE %s"
)

var test bool

func SetTest(prev bool) {
	test = prev
}

var action chan Action

type Action struct {
	symbol *Symbol
	action string
}

type Positions struct {
	Position map[*Symbol]*position
}

type position struct {
	ID              int64
	Symbol          *Symbol
	Shares          int32
	TransactionCost float64
	StrikePrice     float64
	Purchased       time.Time
	Exit            time.Time
	active          bool
	monitored       bool
}

/*func (p *position) Buy(shares int32, price float64) error {

}

func (p *position) Sell(shares int32, price float64) error {

}

func (p *position) Short(shares int32, price float64) error {

}

func (p *position) Cover(shares int32, price float64) error {

}*/

func (p *position) save() error {
	if p.Symbol == nil && p.TransactionCost == 0 && p.StrikePrice == 0 && p.Shares == 0 {
		return NewModelError(IncompletePosition, p)
	}

	if p.ID > 0 { // no need for overwriting
		return nil
	}

	log.Info("saving new position %s", p)
	_, id, err := repository.Exec(fmt.Sprintf(positionInsert, p.Symbol.ID, p.Shares, p.TransactionCost, p.StrikePrice))
	if err != nil {
		return NewModelError(Query, fmt.Sprintf("%s", err))
	}

	p.ID = id

	return nil
}

func (p *position) Load() error {
	if p.ID == 0 {
		if p.Symbol == nil && p.TransactionCost == 0 && p.StrikePrice == 0 && p.Shares == 0 {
			return NewModelError(IncompletePosition, p)
		} else {
			row := repository.QueryRow(fmt.Sprintf(positionSelect, fmt.Sprintf("id=%d", p.ID)))
			p.parseRow(row)

			if p.ID == 0 {
				p.save()
			}

			return nil
		}
	} else {
		return nil
	}
}

func (p *position) Monitor() {

}

func (p *position) Owned() bool {
	return !p.Purchased.IsZero() && p.active && p.Exit.IsZero() && p.Shares > 0
}

func (p *position) String() string {
	return fmt.Sprintf("{ ID: %d, Symbol: %s, Shares: %d, TransactionCost: %f, StrikePrice: %f, purchased: %s, exit: %s, active: %t, monitored: %t }", p.ID, p.Symbol, p.Shares, p.TransactionCost, p.StrikePrice, p.Purchased, p.Exit, p.active, p.monitored)
}

func (p *position) parseRow(row *sql.Row) error {
	return row.Scan(&p.ID, &p.Symbol, &p.Shares, &p.TransactionCost, &p.StrikePrice, &p.Purchased, &p.Exit, &p.active)
}
