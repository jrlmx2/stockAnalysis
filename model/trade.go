package model

import (
	"encoding/xml"
	"fmt"
)

func NewTrade() *Trade { return &Trade{} }

type Trade struct {
	Trade *TradeDetails `xml:"trade"`
}

func (tr *Trade) unmarshal(xmlIn string) (unmarshaler, error) {
	return tr, xml.Unmarshal([]byte(xmlIn), tr)
}

type TradeDetails struct {
	SymbolID              int
	Last                  float32 `xml:"last"`
	Symbol                string  `xml:"symbol"`
	Timestamp             int     `xml:"timestamp"`
	TradedVolume          int64   `xml:"vl"`
	VolumeWeightedAverage float32 `xml:"vwap"`
}

func (td *TradeDetails) Table() string { return "trades" }
func (td *TradeDetails) Data() string {
	return fmt.Sprintf("(NULL,%d,%f,%d,%f,%d,NULL)", td.SymbolID, td.Last, td.TradedVolume, td.VolumeWeightedAverage, td.Timestamp)
}
