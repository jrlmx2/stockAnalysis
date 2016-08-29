package utility

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/jrlmx2/stockAnalysis/utils/oauth"
)

func APIVerison() (*float64, error) {
	req, err := oauthWrapper.Request(uri+"version", "GET")
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n%+v\n", req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	bytes, _ := bufio.NewReader(resp.Body).ReadBytes()
	version := &version{}
	err = xml.Unmarshal(bytes, version)
	if err != nil {
		return nil, err
	}

	t, err := status.Version()
	return &t, err
}

type version struct {
	Response res `xml:"response"`
}

type res struct {
	Version float64 `xml:"version"`
}

func (s *status) Version() float64 {
	return s.Response.Version
}
