package streaming

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"github.com/jrlmx2/stockAnalysis/utils/oauth"
	"github.com/jrlmx2/stockAnalysis/utils/server"
)

const (
	uri = "/market/quotes"
)

type TradeKingStream struct {
	Req  []string
	Resp *http.Response
	S    *bufio.Reader
}

var handler = make(chan *TradeKingStream, 0)

func makeQuery(r []string) string {
	return "?" + strings.Join(r, "&")
}

func OpenStream(r []string) (*TradeKingStream, error) {
	req, err := oauthWrapper.Stream(uri+makeQuery(r), "GET")
	if err != nil {
		fmt.Printf("Openstream failed, trying again. %s", OpenStream(r))
		return nil, err
	}

	fmt.Printf("\n%+v\n", req)
	resp, err := server.Client.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n%+v\n", resp)
	return &TradeKingStream{S: bufio.NewReader(resp.Body), Resp: resp, Req: r},nil
}
