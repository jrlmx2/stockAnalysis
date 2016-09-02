package model

import (
	"database/sql"
	"fmt"

	"github.com/jrlmx2/stockAnalysis/utils/mariadb"
)

const (
	sfindLike     = "select id, symbol from symbols where symbol Like '%%%s%%'"
	sfindOne      = "select id, symbol from symbols where id='%d'"
	sfindSymbols  = "select id, symbol from symbols where symbol in ('%s')"
	sfindSymbol   = "select id, symbol from symbols where symbol = '%s'"
	sfindMany     = "select id, symbol from symbols where id in (%s)"
	sinsert       = "insert into symbols values (NULL,'%s',NULL)"
	sdelete       = "delete from symbols where id=%d"
	sinsertRecord = "(NULL, %s, NULL)"
	qload         = "select * from quotes where symbol_id=%d"
)

type Symbol struct {
	ID         int64
	Symbol     string
	repository *mariadb.Pool
}

func NewSymbol(symbol string) *Symbol { return &Symbol{repository: repository, Symbol: symbol} }

func NewSymbolScan(rows *sql.Rows) (*Symbol, error) {
	s := &Symbol{repository: repository}
	id := &s.ID
	symbol := &s.Symbol

	return s, rows.Scan(id, symbol)
}

func (s *Symbol) Data() string {
	return fmt.Sprintf(sinsertRecord, s.Symbol)
}

func (s *Symbol) Delete() error {
	if s.ID == 0 {
		return NewModelError(NoID)
	}

	_, _, err := s.repository.Exec(fmt.Sprintf(sdelete, s.ID))
	if err != nil {
		return NewModelError(Query)
	}
	return nil
}

func (s *Symbol) Save() error {
	if s.Symbol == "" {
		return NewModelError(NoSymbol)
	}

	if s.ID > 0 { // no need for overwriting
		return nil
	}

	fmt.Printf("%s\n", fmt.Sprintf(sinsert, s.Symbol))
	_, id, err := s.repository.Exec(fmt.Sprintf(sinsert, s.Symbol))
	if err != nil {
		return NewModelError(Query, fmt.Sprintf("%s", err))
	}

	s.ID = id

	return nil
}

func (s *Symbol) Load() error {

	if s.ID == 0 {
		if s.Symbol == "" {
			return NewModelError(EmptySymbol)
		} else {
			row := s.repository.QueryRow(fmt.Sprintf(sfindSymbol, s.Symbol))
			s.parseRow(row)

			if s.ID == 0 {
				s.Save()
			}

			fmt.Printf("Loaded Symbol: %+v\n", s)
			return nil
		}
	} else {
		return nil
	}

}

func (s *Symbol) LoadQuote() *Quote {
	if s.ID == 0 {
		s.Load()
	}

	row := s.repository.QueryRow(fmt.Sprintf(qload, s.ID))

	return ScanQuote(row)
}

func (s *Symbol) LoadTrades() ([]*Trade, error) {
	if s.ID == 0 {
		s.Load()
	}

	rows, err := s.repository.Query(fmt.Sprintf(tfindTrade, s.ID))
	if err != nil {
		return nil, NewModelError(Query, err)
	}
	defer rows.Close()

	return ScanNewTrades(s.Symbol, rows)
}

func (s *Symbol) parseRow(row *sql.Row) error {
	sid := &s.ID
	symbol := &s.Symbol
	return row.Scan(sid, symbol)
}

func (s *Symbol) String() string {
	return fmt.Sprintf("Symbol: %s at %d\n", s.Symbol, s.ID)
}
