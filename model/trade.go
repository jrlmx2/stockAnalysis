package model

import (
	"database/sql"
	"encoding/xml"
	"fmt"

	"github.com/jrlmx2/stockAnalysis/utils/mariadb"
)

const (
	tfindOne    = "select * from trades where id='%d' order by timestamp asc"
	tfindTrades = "select * from trades where symbol_id in ('%s') order by timestamp asc"
	tfindTrade  = "select * from trades where symbol_id = '%d' order by timestamp asc"
	tinsert     = "insert into trades values %s"
	tdelete     = "delete from trades where id=%d"
)

func NewTrade(s *Symbol) *Trade { return &Trade{repository: repository, Symbol: s} }

func ScanNewTrades(s *Symbol, rows *sql.Rows) ([]*Trade, error) {
	defer rows.Close()
	trades := make([]*Trade, 0)
	for rows.Next() {
		t := NewTrade(s)
		err := rows.Scan(&t.Trade.ID, &t.Trade.Last, &t.Trade.Timestamp, &t.Trade.TradedVolume, &t.Trade.VolumeWeightedAverage)
		if err != nil {
			return nil, err
		}
		trades = append(trades, t)
	}

	return trades, nil
}

type Trade struct {
	Trade      *TradeDetails `xml:"trade"`
	Symbol     *Symbol
	repository *mariadb.Pool
}

func (tr *Trade) Unmarshal(xmlIn string) (Unmarshalable, error) {
	return tr, xml.Unmarshal([]byte(xmlIn), tr)
}

type TradeDetails struct {
	ID                    int64
	Last                  float32 `xml:"last"`
	Symbol                string  `xml:"symbol"`
	Timestamp             int     `xml:"timestamp"`
	TradedVolume          int64   `xml:"vl"`
	VolumeWeightedAverage float32 `xml:"vwap"`
}

func (td *Trade) Data() string {
	return fmt.Sprintf("(NULL,%d,%f,%d,%f,%d,NULL)", td.Symbol.ID, td.Trade.Last, td.Trade.TradedVolume, td.Trade.VolumeWeightedAverage, td.Trade.Timestamp)
}

func (t *Trade) Delete() error {
	if t.Trade.ID == 0 {
		return NewModelError(NoTradeID)
	}

	_, _, err := t.repository.Exec(fmt.Sprintf(tdelete, t.Trade.ID))
	if err != nil {
		return NewModelError(QueryError)
	}
	return nil
}

func (t *Trade) Save() error {
	if len(t.Trade.Symbol) < 1 {
		return NewModelError(NoSymbolError)
	}

	t.Symbol = NewSymbol(t.Trade.Symbol)
	err := t.Symbol.Load()
	if err != nil {
		return NewModelError(TradeSaveError, err, t)
	}

	_, _, err = t.repository.Exec(fmt.Sprintf(tinsert, t.Data()))
	if err != nil {
		return NewModelError(QueryError, err)
	}

	return nil
}
