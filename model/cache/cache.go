package cache

import "github.com/jrlmx2/stockAnalysis/model"

/*
* If a symbol is in the cache, its being monitored, therefore, any symbol requested
* to be monitored well be checked in the cache first. If it does not exist in the cache,
* the symbol will be added to a tradeking monitor stream.
 */
type cache struct {
	symbols        []*model.Symbol
	symbolLocation map[string]int
	trades         map[*model.Symbol][]*model.Trade
	quote          map[*model.Symbol]*model.Quote
}

func NewCache() *cache {
	return &cache{
		symbols:        make([]*model.Symbol, 0),
		symbolLocation: make(map[string]int),
		trades:         make(map[*model.Symbol][]*model.Trades, 0),
		quote:          make(map[*model.Symbol]*model.Quote),
	}
}

func (c *cache) Monitor(chan model.Unmarshalable) {

}

func (c *cache) GetSymbol(sym string) *model.Symbol {
	if key, ok := c.symbolLocation[sym]; ok {
		return c.symbols[key]
	}
	return nil
}

func (c *cache) AddSymbol(symbol *model.Symbol) *model.Symbol {
	if sym := c.GetSymbol(symbol.Symbol); sym != nil {
		return sym
	}

	c.symbols = append(c.symbols, symbol)
	c.symbolLocation[symbol.Symbol] = len(c.symbols)
	c.trades[symbol] = make([]*model.Trade, 0)

	return symbol
}
