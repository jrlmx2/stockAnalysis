package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/op/go-logging"

	"github.com/jrlmx2/stockAnalysis/utils/config"
	"github.com/jrlmx2/stockAnalysis/utils/logger"
)

func main() {
	//read config
	conf, file := config.ReadConfig()
	if conf == nil {
		panic("Config -c " + file + " could not be read.")
	}

	//setup logging
	log, err := logger.NewLogger(conf.Logger)
	if err != nil {
		panic(fmt.Sprintf("Logger failed to open with error:%s \n Configuration Details: %+v\n", err, conf.Logger))
	}

	//connect database
	//db, err := mariadb.NewPool(conf.Database)
	if err != nil {
		log.Criticalf("Error opening database at host %s", conf.Database.Host)
		panic(fmt.Sprintf("Error opening database at host %s", conf.Database.Host))
	}

	//establish endpoints
	endpoints := mux.NewRouter()

	//wrap endpoints
	loggedEndpoints := Log(endpoints, log)

	//start server
	http.ListenAndServe(conf.Server.Address, loggedEndpoints)
}

// Log wrapper function for the http server
func Log(handler http.Handler, log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
