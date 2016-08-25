package model

type Stock struct {
	Trades []*Trade
	Symbol *Symbol
	Quotes []*Quote
}

func NewStockWithSymbol(symbol *Symbol) (*Stock, error) {
	trades, err := symbol.LoadTrades()
	if err != nil {
		return nil, err
	}

	return &Stock{Trades: trades, Symbol: symbol, Quotes: make([]*Quote, 0)}, nil
}

func NewStock(s string) (*Stock, error) {
	symbol := NewSymbol(s)

	trades, err := symbol.LoadTrades()
	if err != nil {
		return nil, err
	}

	return &Stock{Trades: trades, Symbol: symbol, Quotes: make([]*Quote, 0)}, nil
}

func (s *Stock) AddTrade(t *Trade) {
	s.Trades = append(s.Trades, t)
}

func (s *Stock) AddQuote(q *Quote) {
	s.Quotes = append(s.Quotes, q)
}
