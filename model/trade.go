package model

import (
	"encoding/xml"
	"time"

	"github.com/jrlmx2/stockAnalysis/utils/influxdb"
)

const (
	tfindOne    = "select * from trades where id='%d' order by timestamp asc"
	tfindTrades = "select * from trades where symbol_id in ('%s') order by timestamp asc"
	tfindTrade  = "select id, symbol_id, last, timestamp, tradedvolume, vwa, created_at from trades where symbol_id = '%d' order by timestamp asc"
	tinsert     = "insert into trades values %s"
	tdelete     = "delete from trades where id=%d"
)

func NewTrade(s string) *Trade { return &Trade{Symbol: s} }

func NewTradeU() *Trade { return &Trade{} }

/*func ScanNewTrades(s string, rows *sql.Rows) ([]*Trade, error) {
	defer rows.Close()
	trades := make([]*Trade, 0)
	for rows.Next() {
		t := NewTrade(s)
		err := rows.Scan(&t.ID, &t.SymbolID, &t.Last, &t.Timestamp, &t.TradedVolume, &t.VolumeWeightedAverage, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		trades = append(trades, t)
	}

	return trades, nil
}*/

type Trade struct {
	Trade                 xml.Name `xml:"trade"`
	ID                    int64
	SymbolID              int64
	Last                  float64 `xml:"last"`
	Symbol                string  `xml:"symbol"`
	Timestamp             int     `xml:"timestamp"`
	TradedVolume          int64   `xml:"vl"`
	VolumeWeightedAverage float64 `xml:"vwap"`
	CreatedAt             time.Time
}

func (t *Trade) Unmarshal(xmlIn string) (Unmarshalable, error) {
	return t, xml.Unmarshal([]byte(xmlIn), t)
}

func (td *Trade) Labels() map[string]interface{} {

	tags := make(map[string]interface{})

	tags["vwa"] = td.VolumeWeightedAverage
	tags["last"] = td.Last
	tags["volume"] = td.TradedVolume

	return tags
}

func (td *Trade) Tags() map[string]string {
	return map[string]string{}
}

/*func (t *Trade) Delete() error {
	if t.ID == 0 {
		return NewModelError(NoTradeID)
	}

	_, _, err := t.repository.Exec(fmt.Sprintf(tdelete, t.ID))
	if err != nil {
		return NewModelError(Query)
	}
	return nil
}*/

func (t *Trade) Save() error {

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

	return influxdb.AddPoint("tradeking_"+t.Symbol, "stocks", t.Tags(), t.Labels())
}

/*func (td *Trade) Data() string {
	return fmt.Sprintf("(NULL,%f,%d,%d,%d,%f,NULL)", td.Last, td.SymbolID, td.Timestamp, td.TradedVolume, td.VolumeWeightedAverage)
}*/
