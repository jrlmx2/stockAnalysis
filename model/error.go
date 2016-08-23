package model

import "fmt"

const (
	NoIDError     ModelError = "No ID was passed in when to find"
	QueryError    ModelError = "SQL query failed with error %s"
	NoSymbolError ModelError = "Symbol not present on save"
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
