package model

import (
	"fmt"
)

const (
	NoID               ModelError = "No ID was passed in when to find"
	Query              ModelError = "SQL query failed with error %s"
	NoSymbol           ModelError = "Symbol not present on save"
	EmptySymbol        ModelError = "There is no data to save. Symbol is empty."
	NoTradeID          ModelError = "The current trade does not have an ID associated with it."
	TradeSave          ModelError = "Trade saving errored: %s, trade: %s"
	QuoteSave          ModelError = "Quote saving failed with %s, quote: %s"
	NoName             ModelError = "Watchlist has no name"
	UniqueName         ModelError = "Watchlist name is already taken"
	Load               ModelError = "%s loading error %s"
	IncompletePosition ModelError = "Incomplete postion, cannot work with %+v"
)

type ModelError string

func (me *ModelError) toString() string {
	return fmt.Sprintf("%v", me)
}

type ModelErrorContainer struct {
	err  *ModelError
	args []interface{}
}

func (mec *ModelErrorContainer) Error() string {
	return fmt.Sprintf(mec.err.toString(), mec.args...)
}

func NewModelError(err ModelError, args ...interface{}) *ModelErrorContainer {
	return &ModelErrorContainer{err: &err, args: args}
}
