package model

type Stock struct {
	Trades []*Trade
	Symbol *Symbol
	Quotes *Quote
}

func NewStockWithSymbol(symbol *Symbol) (*Stock, error) {
	return newStock(symbol)
}

func NewStock(s string) (*Stock, error) {
	symbol := NewSymbol(s)
	return newStock(symbol)
}

func newStock(symbol *Symbol) (*Stock, error) {
	trades, err := symbol.LoadTrades()
	if err != nil {
		return nil, err
	}

	quote := symbol.LoadQuote()

	return &Stock{Trades: trades, Symbol: symbol, Quotes: quote}, nil
}

func (s *Stock) AddTrade(t *Trade) {
	s.Trades = append(s.Trades, t)
}
