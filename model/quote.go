package model

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/jrlmx2/stockAnalysis/utils/mariadb"
)

const (
	quoteDateFormat = "2006-01-02T15:04:05-07:00"
	readableDate    = "01 Jan 2006 at 15:04:05 MST"
	qinsert         = "insert into quotes values %s"
)

func NewQuote(s *Symbol) *Quote { return &Quote{repository: repository, Symbol: s} }

type Quote struct {
	repository *mariadb.Pool
	Quote      *QuoteDetails `xml:"quote"`
	Symbol     *Symbol
}

func NewQuoteU() *Quote { return &Quote{repository: repository} }

func (unm *Quote) Unmarshal(xmlIn string) (Unmarshalable, error) {
	return unm, xml.Unmarshal([]byte(xmlIn), unm)
}

type QuoteDetails struct {
	ID        int64
	Ask       float32 `xml:"ask"`
	AskVolume int     `xml:"asksz"`
	Bid       float32 `xml:"bid"`
	BidVolume int     `xml:"bidsz"`
	Datetime  string  `xml:"datetime"`
	Symbol    string  `xml:"symbol"`
	Timestamp int     `xml:"timestamp"`
}

func (q *Quote) Data() string {
	return fmt.Sprintf("(NULL, %f, %d, %f, %d, %d, %d)", q.Quote.Ask, q.Quote.AskVolume, q.Quote.Bid, q.Quote.BidVolume, q.Symbol.ID, q.Quote.Timestamp)
}

func (q *Quote) Date() (string, error) {
	t, err := time.Parse(quoteDateFormat, q.Quote.Datetime)
	if err != nil {
		return "", err
	}

	return t.Format(readableDate), nil
}

func (q *Quote) Save() error {
	if q.Symbol == nil && q.Quote.Symbol == "" {
		return NewModelError(EmptySymbol)
	}

	if q.Symbol == nil {
		q.Symbol = NewSymbol(q.Quote.Symbol)
		err := q.Symbol.Load()
		if err != nil {
			return NewModelError(QuoteSave, err, q)
		}
	}

	_, id, err := q.repository.Exec(fmt.Sprintf(qinsert, q.Data()))
	if err != nil {
		return NewModelError(Query, err)
	}

	q.Quote.ID = id

	return nil
}
