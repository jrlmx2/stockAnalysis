package streaming

import (
	"fmt"
	"strings"
)

type unmarshaler interface {
	unmarshal(string) (unmarshaler, error)
}

func umarshal(xml string) (unmarshaler, error) {
	if strings.Contains(xml, "quote") {
		return NewQuote().unmarshal(xml)
	}

	if strings.Contains(xml, "trade") {
		return NewTrade().unmarshal(xml)
	}

	fmt.Printf("XML not identified %s", xml)
	return NewQuote(), nil
}
