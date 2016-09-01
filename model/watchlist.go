package model

import (
	"fmt"

	"github.com/jrlmx2/stockAnalysis/utils/mariadb"
)

type WatchList struct {
	Symbols    []*Symbol
	Stocks     []*Stock
	Options    *Options
	Name       string
	ID         int64
	repository *mariadb.Pool
}

const (
	winsert      = "insert into watchlist values %s"
	wfind        = "select * from watchlist where name='%s'"
	wfindsymbols = "select s.symbol_id, s.symbol from watchlist_symbols ws join symbols s on s.id = ws.symbol_id where watchlist_id = '%d'"
)

func (w *WatchList) GetSymbols() error {
	if w.ID == 0 {
		err := w.Load()
		if err != nil {
			return err
		}
	}

	rows, err := w.repository.Query(fmt.Sprintf(wfindsymbols, w.ID))
	if err != nil {
		return NewModelError(Query, err)
	}
	defer rows.Close()

	symbols := make([]*Symbol, 0)
	stocks := make([]*Stock, 0)
	for rows.Next() {
		s, err := NewSymbolScan(rows)
		if err != nil {
			return err
		}
		symbols = append(symbols, s)

		stock, err := NewStockWithSymbol(s)
		if err != nil {
			return err
		}
		stocks = append(stocks, stock)
	}

	w.Symbols = symbols
	w.Stocks = stocks

	return nil
}

func NewWatchList(name string) (*WatchList, error) {
	list := &WatchList{Name: name, repository: repository}
	list.Save()
	return list, list.Save()
}

func (w *WatchList) LoadList(list string) error {
	if list == "" {
		return NewModelError(NoName)
	}

	row := w.repository.QueryRow(fmt.Sprint(wfind, list))

	name := &w.Name
	id := &w.ID
	row.Scan(id, name)
	return nil
}

func (w *WatchList) Load() error {
	if w.Name == "" {
		return NewModelError(NoName)
	}

	row := w.repository.QueryRow(fmt.Sprint(wfind, w.Name))

	name := &w.Name
	id := &w.ID
	row.Scan(id, name)
	return nil
}

func (w *WatchList) Save() error {
	if w.Name == "" {
		return NewModelError(NoName)
	}

	err := w.Load()
	if err != nil {
		return NewModelError(Load, "Watchlist", err)
	}
	if w.ID != 0 {
		return nil
	}

	_, id, err := w.repository.Exec(fmt.Sprintf(winsert, w.Name))
	if err != nil {
		return NewModelError(Query, err)
	}

	w.ID = id

	return nil
}

type Options struct {
	//Indicators []*Indicator
}
