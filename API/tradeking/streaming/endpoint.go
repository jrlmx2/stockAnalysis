package streaming

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"github.com/jrlmx2/stockAnalysis/utils/oauth"
)

const (
	uri = "market/quotes"
)

func makeQuery(r []string) string {
	return "?" + strings.Join(r, "&")
}

func OpenStream(r []string, handler chan<- bufio.Reader) error {
	req, err := oauthWrapper.Stream(uri+makeQuery(r), "GET")
	if err != nil {
		return err
	}
	fmt.Printf("\n%+v\n", req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("\n%+v\n", resp)

	handler <- bufio.NewReader(resp.Body)
	return nil
}
