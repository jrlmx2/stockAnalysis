package model

import "github.com/jrlmx2/stockAnalysis/utils/mariadb"

type WatchList struct {
	Symbols    []*Symbol
	Stocks     []*Stock
	Options    *Options
	repository *mariadb.Pool
}

func NewWatchList(symbols []*Symbol) (*WatchList, error) {

	stocks := make([]*Stock, 0)
	for _, symbol := range symbols {
		stock, err := NewStockWithSymbol(symbol)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}
	return &WatchList{repository: repository, Symbols: symbols, Stocks: stocks, Options: &Options{}}, nil
}

type Options struct {
	//Indicators []*Indicator
}
