package model

import "fmt"

const (
	NoIDError        ModelError = "No ID was passed in when to find"
	QueryError       ModelError = "SQL query failed with error %s"
	NoSymbolError    ModelError = "Symbol not present on save"
	EmptySymbolError ModelError = "There is no data to save. Symbol is empty."
	NoTradeID        ModelError = "The current trade does not have an ID associated with it."
	TradeSaveError   ModelError = "Trade savinging errored: %s, trade: %s"
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
