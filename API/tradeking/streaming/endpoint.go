package streaming

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/jrlmx2/stockAnalysis/utils/oauth"
	"github.com/jrlmx2/stockAnalysis/utils/server"
)

const (
	uri = "market/quotes"
)

var handler = make(chan *bufio.Reader, 0)

func makeQuery(r []string) string {
	return "?" + strings.Join(r, "&")
}

func OpenStream(r []string) error {
	req, err := oauthWrapper.Stream(uri+makeQuery(r), "GET")
	if err != nil {
		return err
	}
	fmt.Printf("\n%+v\n", req)
	resp, err := server.Client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("\n%+v\n", resp)

	handler <- bufio.NewReader(resp.Body)
	return nil
}
