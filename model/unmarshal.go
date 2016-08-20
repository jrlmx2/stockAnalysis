package model

type Unmarshalable interface {
	Unmarshal(string) (unmarshaler, error)
}
