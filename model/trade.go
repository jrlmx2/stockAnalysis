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

func NewTrade(s string) *Trade { return &Trade{repository: repository, Symbol: s} }

func NewTradeU() *Trade { return &Trade{repository: repository} }

func ScanNewTrades(s string, rows *sql.Rows) ([]*Trade, error) {
	defer rows.Close()
	trades := make([]*Trade, 0)
	for rows.Next() {
		t := NewTrade(s)
		err := rows.Scan(&t.ID, &t.Last, &t.Timestamp, &t.TradedVolume, &t.VolumeWeightedAverage)
		if err != nil {
			return nil, err
		}
		trades = append(trades, t)
	}

	return trades, nil
}

type Trade struct {
	Trade                 xml.Name `xml:"trade"`
	ID                    int64
	SymbolID              int64
	Last                  float32 `xml:"last"`
	Symbol                string  `xml:"symbol"`
	Timestamp             int     `xml:"timestamp"`
	TradedVolume          int64   `xml:"vl"`
	VolumeWeightedAverage float32 `xml:"vwap"`
	repository            *mariadb.Pool
}

func (t *Trade) Unmarshal(xmlIn string) (Unmarshalable, error) {
	return t, xml.Unmarshal([]byte(xmlIn), t)
}

func (td *Trade) Data() string {
	return fmt.Sprintf("(NULL,%f,%d,%d,%d,%f,NULL)", td.Last, td.SymbolID, td.Timestamp, td.TradedVolume, td.VolumeWeightedAverage)
}

func (t *Trade) Delete() error {
	if t.ID == 0 {
		return NewModelError(NoTradeID)
	}

	_, _, err := t.repository.Exec(fmt.Sprintf(tdelete, t.ID))
	if err != nil {
		return NewModelError(Query)
	}
	return nil
}

func (t *Trade) Save() error {
	fmt.Printf("%+v", t)
	if len(t.Symbol) < 1 {
		return NewModelError(NoSymbol)
	}

	if t.SymbolID == 0 {
		Symbol := NewSymbol(t.Symbol)
		err := Symbol.Load()
		if err != nil {
			return NewModelError(TradeSave, err, t)
		}
		t.SymbolID = Symbol.ID
	}

	_, id, err := t.repository.Exec(fmt.Sprintf(tinsert, t.Data()))
	if err != nil {
		fmt.Printf("No Query Err %+v", t)
		return NewModelError(Query, err)
	}

	t.ID = id

	return nil
}
