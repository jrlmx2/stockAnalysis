package streaming

import "encoding/xml"

func NewTrade() *Trade { return &Trade{} }

type Trade struct {
	Trade *TradeDetails `xml:"trade"`
}

func (tr *Trade) unmarshal(xmlIn string) (unmarshaler, error) {
	return tr, xml.Unmarshal([]byte(xmlIn), tr)
}

type TradeDetails struct {
	Last      float32 `xml:"last"`
	Symbol    string  `xml:"symbol"`
	Timestamp int     `xml:"timestamp"`
	Amount    int     `xml:"vl"`
	Vwap      float32 `xml:"vwap"`
}
