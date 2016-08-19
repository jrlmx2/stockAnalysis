package streaming

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"github.com/jrlmx2/stockAnalysis/lib/oauth"
)

const (
	uri = "market/quotes"
)

func makeQuery(r []string) string {
	return "?" + strings.Join(r, "&")
}

func OpenStream(r []string, handler chan<- string) error {
	req, err := oauthWrapper.Stream(uri+makeQuery(r), "GET")
	if err != nil {
		return nil
	}
	fmt.Printf("\n%+v\n", req)
	resp, err := http.DefaultClient.Do(req)
	fmt.Printf("\n%+v\n", resp)

	reader := bufio.NewReader(resp.Body)
	content := ""
	for {
		line, err := reader.ReadString('>')
		if err != nil {
			return err
		}

		sline := string(line)

		if strings.Contains(sline, "/") && strings.Contains(sline, "status") {
			content = ""
			continue
		}

		if strings.Contains(sline, "/") && (strings.Contains(sline, "quote") || strings.Contains(sline, "trade")) {
			content += sline
			fmt.Printf("CONTENT: %+v\n\n", content)
			handler <- content
			content = ""
		} else {
			content += sline
		}
	}

}

func unmarshal(xml string) {

}
