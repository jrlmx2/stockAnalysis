package streaming

import (
	"bufio"
	"net/http"
)

var StreamInput chan interface{}

type TradeKingStream struct {
	Req  []string
	Resp *http.Response
	S    *bufio.Reader
}

func (tks *TradeKingStream) Connection() *bufio.Reader {
	return tks.S
}

func (tks *TradeKingStream) Reopen() {
	OpenStream(tks.Req)
}
