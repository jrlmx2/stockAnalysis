package watchlist

import (
	"github.com/gorilla/mux"
	"github.com/jrlmx2/stockAnalysis/API/homegrown/watchlist"
)

const watchlist_base_uri = "watchlist"

func Endpoints(handler *mux.Router) *mux.Router {
	watchlist.Endpoints(handler)

	return handler
}
