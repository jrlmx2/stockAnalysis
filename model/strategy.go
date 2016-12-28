package model

type strategy struct {
	//rules     []*rule
	trade bool
	//	account   *accountManager
	positions *position
}

func NewStrategy(trade bool) (*strategy, error) {
	return &strategy{}, nil
}
