package sec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Result struct {
		Totalrows int `json:"totalrows"`
		Rows      []struct {
			Rownum int `json:"rownum"`
			Values []struct {
				Field string `json:"field"`
				Value string `json:"value"`
			} `json:"values"`
		} `json:"rows"`
	} `json:"result"`
}

func Get(resource, filter string) (*Response, error) {

	response, err := http.Get(fmt.Sprintf(uri, resource, filter, key))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(body, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil

}
