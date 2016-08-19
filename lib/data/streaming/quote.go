package streaming

import (
	"encoding/xml"
	"time"
)

const (
	quoteDateFormat = "2006-01-02T15:04:05-07:00"
	readableDate    = "01 Jan 2006 at 15:04:05 MST"
)

func NewQuote() *Quote { return &Quote{} }

type Quote struct {
	Quote QuoteDetails `xml:"quote"`
}

func (unm *Quote) unmarshal(xmlIn string) (unmarshaler, error) {
	return unm, xml.Unmarshal([]byte(xmlIn), unm)
}

type QuoteDetails struct {
	Ask       float32 `xml:"ask"`
	AskVolume int     `xml:"asksz"`
	Bid       float32 `xml:"bid"`
	BidVolume int     `xml:"bidsz"`
	Datetime  string  `xml:"datetime"`
	Symbol    string  `xml:"symbol"`
	Timestamp int     `xml:"timestamp"`
}

func (q *Quote) Date() (string, error) {
	t, err := time.Parse(quoteDateFormat, q.Quote.Datetime)
	if err != nil {
		return "", err
	}

	return t.Format(readableDate), nil
}
