package model

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/jrlmx2/stockAnalysis/utils/mariadb"
)

const (
	findLike     = "select id, symbol from symbols where symbol Like '%%%s%%'"
	findOne      = "select id, symbol from symbols where id='%d'"
	findSymbols  = "select id, symbol from symbols where symbol in ('%s')"
	findMany     = "select id, symbol from symbols where id in (%s)"
	insert       = "insert into symbols values (NULL, '%s', NULL)"
	delete       = "delete from symbols where id=%d"
	insertRecord = "(NULL, %s, NULL)"
)

type Symbol struct {
	ID         int64
	Symbol     string
	repository *mariadb.Pool
}

type SymbolCollection struct {
	Symbols    []*Symbol
	repository *mariadb.Pool
}

func NewSymbolCollection() *SymbolCollection { return &SymbolCollection{repository: repository} }
func NewSymbol() *Symbol                     { return &Symbol{repository: repository} }
func NewSymbolScan() (*Symbol, *int64, *string) {
	s := &Symbol{repository: repository}
	return s, &s.ID, &s.Symbol
}

func (s *Symbol) Data() string {
	return fmt.Sprintf(insertRecord, s.Symbol)
}

func (s *Symbol) Delete() error {
	if s.ID == 0 {
		return NewModelError(NoIDError)
	}

	_, _, err := s.repository.Exec(fmt.Sprintf(delete, s.ID))
	if err != nil {
		return NewModelError(QueryError)
	}
	return nil
}

func (s *Symbol) Save() error {
	if s.Symbol == "" {
		return NewModelError(NoSymbolError)
	}

	fmt.Printf("%s", fmt.Sprintf(insert, s.Symbol))
	_, id, err := s.repository.Exec(fmt.Sprintf(insert, s.Symbol))
	if err != nil {
		return NewModelError(QueryError, fmt.Sprintf("%s", err))
	}

	s.ID = id

	return nil
}

func (s *SymbolCollection) SaveAll() error {

	syms := make([]string, 0)
	for _, sym := range s.Symbols {
		if sym.ID > 0 {
			continue
		}

		syms = append(syms, sym.Data())
	}

	if len(syms) == 0 {
		return nil
	}

	_, _, err := s.repository.Exec(fmt.Sprintf(insert, strings.Join(syms, ",")))
	if err != nil {
		return NewModelError(QueryError, fmt.Sprintf("%s", err))
	}

	s.FindSymbols()

	return nil
}
func (s *SymbolCollection) FindLike(symbol string) error {

	rows, err := s.repository.Query(fmt.Sprintf(findLike, symbol))
	if err != nil {
		return NewModelError(QueryError, fmt.Sprintf("%s", err))
	}

	s.parseRows(rows)

	return nil
}

func (s *SymbolCollection) FindSymbols() error {
	syms := make([]string, 0)
	for _, sym := range s.Symbols {
		if sym.ID > 0 {
			continue
		}

		syms = append(syms, sym.Symbol)
	}

	if len(syms) == 0 {
		return nil
	}

	rows, err := s.repository.Query(fmt.Sprintf(findSymbols, strings.Join(syms, ",")))
	if err != nil {
		return NewModelError(QueryError, fmt.Sprintf("%s", err))
	}

	s.parseRows(rows)

	return nil

}

func (s *Symbol) Load(id int64) error {
	if id == 0 {
		return NewModelError(NoIDError)
	}

	row := s.repository.QueryRow(fmt.Sprintf(findOne, id))
	sid := &s.ID
	symbol := &s.Symbol
	return row.Scan(sid, symbol)
}

func (s *SymbolCollection) LoadMany(ids []int64) error {
	if ids == nil {
		fmt.Println("No IDs passed in")
		return nil
	}

	sids := make([]string, 0)
	for _, value := range ids {
		s := strconv.FormatInt(value, 10)
		sids = append(sids, s)
	}

	rows, err := s.repository.Query(fmt.Sprintf(findMany, "'"+strings.Join(sids, "','")+"'"))
	if err != nil {
		return err
	}

	s.parseRows(rows)

	return nil
}

func (s *SymbolCollection) parseRows(rows *sql.Rows) {

	for rows.Next() {
		symbol, id, name := NewSymbolScan()
		rows.Scan(id, name)
		s.Symbols = append(s.Symbols, symbol)
	}
}

func (s *SymbolCollection) String() string {
	result := fmt.Sprintf("SymbolCollection has the following symbols: \n")
	for _, value := range s.Symbols {
		result += value.String()
	}
	result += "\n\n"
	return result
}

func (s *Symbol) String() string {
	return fmt.Sprintf("Symbol: %s at %d\n", s.Symbol, s.ID)
}
