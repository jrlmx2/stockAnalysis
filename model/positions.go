package model

import (
	"time"
)

var preview = true

func SetPreview(prev bool) {
	preview = prev
}

type Positions struct {
	Position map[*Symbol]*position
}

type position struct {
	ID               int64
	Shares           int64
	TransactionCoast float64
	Value            float64
	Purchased        time.Time
}

/*func (p *position) Buy() error {

}

func (p *position) Sell() error {

}

func (p *position) Short() error {

}

func (p *position) Cover() error {

}*/
