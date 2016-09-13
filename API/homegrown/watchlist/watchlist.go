package watchlist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

const watchlist_base_uri = "watchlist"

func Endpoints(handler *mux.Router) {
	handler.HandleFunc(watchlist_base_uri, Update).Methods("POST")
	handler.HandleFunc(watchlist_base_uri, Find).Methods("GET")
}

func Find(w http.ResponseWriter, r *http.Request) {

}

func Update(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("\n\n Parsing body of request: %+v failed with %s\n\n", r, err)
		w.WriteHeader(403)
		w.Write(fmt.Sprintf("\n\n Parsing body of request: %+v failed with %s\n\n", r, err))
		return
	}

	var requestBody Request
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		fmt.Printf("\n\n Unmarshaling %s failed with %s\n\n", body, err)
		w.WriteHeader(403)
		w.Write(fmt.Sprintf("\n\n Unmarshaling %s failed with %s\n\n", body, err))
		return
	}
	list, err := model.LoadList(requestBody.List)
	if err != nil || list.ID == 0 {
		list, err = model.NewList(requestBody.List)
	}

	if requestBody.Symbols {
		err := list.UpdateSymbols(reqestBody.Symbols)
		if err != nil {
			fmt.Printf("\n\nerrrororororoorr %s\n\n", err)
			return
		}
	}

	w.Header("Status", "200")
}

type Request struct {
	List    string   `json:"list"`
	Symbols []string `json:"symbols,omitempty"`
	Options []string `json:"options,omitempty"`
}
