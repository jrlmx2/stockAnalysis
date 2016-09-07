package model

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"time"

	"github.com/jrlmx2/stockAnalysis/utils/mariadb"
)

const (
	quoteDateFormat = "2006-01-02T15:04:05-07:00"
	readableDate    = "01 Jan 2006 at 15:04:05 MST"
	qinsert         = "insert into quotes values %s on duplicate key update %s"
)

func NewQuote(s string) *Quote { return &Quote{repository: repository, Symbol: s} }

type Quote struct {
	repository *mariadb.Pool
	Quote      xml.Name `xml:"quote"`
	ID         int64
	SymbolID   int64
	Ask        float32 `xml:"ask"`
	AskVolume  int     `xml:"asksz"`
	Bid        float32 `xml:"bid"`
	BidVolume  int     `xml:"bidsz"`
	Datetime   string  `xml:"datetime"`
	Symbol     string  `xml:"symbol"`
	Timestamp  int     `xml:"timestamp"`
}

func NewQuoteU() *Quote { return &Quote{repository: repository} }

func ScanQuote(row *sql.Row) *Quote {
	quote := NewQuoteU()
	row.Scan(&quote.ID, &quote.Ask, &quote.AskVolume, &quote.Bid, &quote.BidVolume, &quote.SymbolID, &quote.Timestamp)
	return quote
}

func (unm *Quote) Unmarshal(xmlIn string) (Unmarshalable, error) {
	return unm, xml.Unmarshal([]byte(xmlIn), unm)
}

func (q *Quote) Data() (string, string) {
	return fmt.Sprintf("(NULL, %f, %d, %f, %d, %d, %d)", q.Ask, q.AskVolume, q.Bid, q.BidVolume, q.SymbolID, q.Timestamp), fmt.Sprintf(" ask=%f, askvolume=%d, bid=%f, bidvolume=%d, timestamp=%d", q.Ask, q.AskVolume, q.Bid, q.BidVolume, q.Timestamp)
}

func (q *Quote) Date() (string, error) {
	t, err := time.Parse(quoteDateFormat, q.Datetime)
	if err != nil {
		return "", err
	}

	return t.Format(readableDate), nil
}

func (q *Quote) Save() error {
	if q.Symbol == "" {
		return NewModelError(EmptySymbol)
	}

	if q.SymbolID == 0 {
		Symbol := NewSymbol(q.Symbol)
		err := Symbol.Load()
		if err != nil {
			return NewModelError(QuoteSave, err, q)
		}
		q.SymbolID = Symbol.ID
	}

	insData, updateData := q.Data()
	_, id, err := q.repository.Exec(fmt.Sprintf(qinsert, insData, updateData))
	if err != nil {
		return NewModelError(Query, err)
	}

	q.ID = id

	return nil
}
