package model

import (
	"fmt"
	"strings"

	"github.com/jrlmx2/stockAnalysis/utils/mariadb"
)

type WatchList struct {
	Symbols    []*Symbol
	Stocks     map[Symbol]*Stock
	Options    *Options
	Name       string
	ID         int64
	repository *mariadb.Pool
}

const (
	winsert        = "insert into watchlist values %s"
	wdeletesymbols = "delete from watchlist_symbols where symbol_id=%d"
	winsertsymbols = "insert into watchlist_symbols values %s"
	wfind          = "select * from watchlist where name='%s'"
	wfindsymbols   = "select s.symbol_id, s.symbol from watchlist_symbols ws join symbols s on s.id = ws.symbol_id where watchlist_id = '%d'"
)

func (w *WatchList) SymbolData() string {
	inserts := make([]string, 0)
	for _, symbol := range w.Symbols {
		inserts = append(inserts, fmt.Sprintf("(NULL,%d,%d)", symbol.ID, w.ID))
	}
	return strings.Join(inserts, ",")
}

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

	symbols := make([]*Symbol, 0)
	stocks := make(map[Symbol]*Stock, 0)
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
		stocks[*s] = stock
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

	w.repository.QueryRow(fmt.Sprint(wfind, list)).Scan(&w.ID, &w.Name)

	w.GetSymbols()
	return nil
}

func (w *WatchList) Load() error {
	if w.Name == "" {
		return NewModelError(NoName)
	}

	w.repository.QueryRow(fmt.Sprint(wfind, w.Name)).Scan(&w.ID, &w.Name)

	w.GetSymbols()
	return nil
}

func (w *WatchList) UpdateSymbols(symbols []string) error {
	if w.ID == 0 {
		w.Load()
	}

	newSymbolList := make([]*Symbol, 0)
	for _, symbol := range symbols {
		s := NewSymbol(symbol)
		s.Load()
		newSymbolList = append(newSymbolList, s)
	}

	for _, symbol := range w.Symbols {
		if !contains(newSymbolList, symbol) {
			delete(w.Stocks, *symbol)
		}
	}

	w.Symbols = newSymbolList

	w.Save()

	return w.GetStocks()
}

func (w *WatchList) GetStocks() error {
	stocks := make(map[Symbol]*Stock, 0)
	for _, s := range w.Symbols {
		stock, err := NewStockWithSymbol(s)
		if err != nil {
			return err
		}
		stocks[*s] = stock
	}
	return nil
}

func contains(contents []*Symbol, test *Symbol) bool {
	for _, val := range contents {
		if val.Symbol == test.Symbol {
			return true
		}
	}
	return false
}

func (w *WatchList) exists(name string) bool {
	for _, symbol := range w.Symbols {
		if symbol.Symbol == name {
			return true
		}
	}

	return false
}

func (w *WatchList) Save() error {
	if w.Name == "" {
		return NewModelError(NoName)
	}

	if w.ID == 0 {
		err := w.Load()

		if err != nil {
			return NewModelError(Load, "Watchlist", err)
		}

		_, id, err := w.repository.Exec(fmt.Sprintf(winsert, w.Name))
		if err != nil {
			return NewModelError(Query, err)
		}

		w.ID = id
	}

	if w.Symbols != nil {
		_, _, err := w.repository.Exec(fmt.Sprintf(winsertsymbols, w.SymbolData()))
		if err != nil {
			return err
		}
	}

	//eventually options

	return nil
}

type Options struct {
	//Indicators []*Indicator
}
