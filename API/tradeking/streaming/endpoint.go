package streaming

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/jrlmx2/stockAnalysis/utils/oauth"
	"github.com/jrlmx2/stockAnalysis/utils/server"
)

const (
	uri = "/market/quotes"
)

func makeQuery(r []string) string {
	return "?" + strings.Join(r, "&")
}

func OpenStream(r []string) {
	req, err := oauthWrapper.Stream(uri+makeQuery(r), "GET")
	if err != nil {
		OpenStream(r)
		fmt.Printf("Openstream failed, trying again")
	}

	fmt.Printf("\n%+v\n", req)
	resp, err := server.Client.Do(req)
	if err != nil {
		OpenStream(r)
		fmt.Printf("Error running the http request")
	}

	fmt.Printf("\n%+v\n", resp)
	StreamInput <- &TradeKingStream{S: bufio.NewReader(resp.Body), Resp: resp, Req: r}
}
