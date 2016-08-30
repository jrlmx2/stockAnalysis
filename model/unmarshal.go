package model

type Unmarshalable interface {
	Unmarshal(string) (Unmarshalable, error)
	Save() error
}
