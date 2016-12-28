package tradeking

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jrlmx2/stockAnalysis/API/tradeking/streaming"
	"github.com/jrlmx2/stockAnalysis/utils/config"
	"github.com/jrlmx2/stockAnalysis/utils/oauth"
)

// EstablishEndpoints is used for appending tradeking public API calls to the Server
func EstablishEndpoints(handler *mux.Router, streams chan interface{}) *mux.Router {
	conf := config.ReadConfigPath("API/tradeking/api.conf")

	//logger, _ := logger.NewLogger(conf.Logger)

	oauthWrapper.SetCredentials(conf.API["tradeking"].OAuthToken, conf.API["tradeking"].OAuthSecret)
	oauthWrapper.SetClient(conf.API["tradeking"].Key, conf.API["tradeking"].Secret)

	streaming.StreamInput = streams

	return Endpoints(handler)

}

func StreamOpener(w http.ResponseWriter, r *http.Request) {
	symbols := fmt.Sprintf("symbols=%s", r.Form["symbols"])

	streaming.OpenStream([]string{symbols})

	w.Write([]byte("Done."))

}

func Endpoints(handler *mux.Router) *mux.Router {
	handler.HandleFunc("/stream", StreamOpener).Methods("GET")

	return handler
}
