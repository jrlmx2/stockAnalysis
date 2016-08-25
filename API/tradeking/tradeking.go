package tradeking

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/jrlmx2/stockAnalysis/API/tradeking/streaming"
	"github.com/jrlmx2/stockAnalysis/utils/config"
	"github.com/jrlmx2/stockAnalysis/utils/logger"
)

// EstablishEndpoints is used for appending tradeking public API calls to the Server
func EstablishEndpoints(handler http.Handler) http.Handler {
	conf := config.ReadConfigPath("./api_config")
	fmt.Printf("\n\n%+v\n\n", conf)

	logger, _ := logger.NewLogger(conf.Logger)
	fmt.Printf("%+v", logger)

	oauthWrapper.SetCredentials(conf.API["tradeking"].OAuthToken, conf.API["tradeking"].OAuthSecret)
	oauthWrapper.SetClient(conf.API["tradeking"].Key, conf.API["tradeking"].Secret)

	streams := streaming.ProcessStreams(logger)
	streams = nil // do somethign with streams

}

func StreamOpener(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	symbols := strings.Join(r.Form["symbols"], ",")

}

func Endpoints(handler http.Handler) http.Handler {
	handler.HandleFunc("/stream", StreamOpener).Methods("GET")

	return handler
}
