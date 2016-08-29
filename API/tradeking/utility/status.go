package utility

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"github.com/jrlmx2/stockAnalysis/utils/oauth"
)

const statusTimeFormat = "Mon, 02 Jan 2006 15:04:05 MST"

func ServerStatus() (*time.Time, error) {
	req, err := oauthWrapper.Request(uri+"status", "GET")
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n%+v\n", req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	bytes, _ := bufio.NewReader(resp.Body).ReadBytes()
	status := &status{}
	err = xml.Unmarshal(bytes, status)
	if err != nil {
		return nil, err
	}

	t, err := status.Time()
	return &t, err
}

type status struct {
	Response res `xml:"response"`
}

type res struct {
	Date string `xml:"time"`
}

func (s *status) Time() (time.Time, error) {
	return time.Parse(statusTimeFormat, s.Response.Date)
}
