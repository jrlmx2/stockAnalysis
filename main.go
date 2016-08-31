package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	logging "github.com/op/go-logging"

	"github.com/jrlmx2/stockAnalysis/API/tradeking"
	"github.com/jrlmx2/stockAnalysis/model"
	"github.com/jrlmx2/stockAnalysis/utils/config"
	"github.com/jrlmx2/stockAnalysis/utils/logger"
	"github.com/jrlmx2/stockAnalysis/utils/mariadb"
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
	db, err := mariadb.NewPool(conf.Database)
	if err != nil {
		log.Criticalf("Error opening database at host %s", conf.Database.Host)
		panic(fmt.Sprintf("Error opening database at host %s", conf.Database.Host))
	}

	model.SetRepository(db)

	log.Info("Database connected")
	//establish endpoints
	endpoints := mux.NewRouter()
	endpoints = tradeking.EstablishEndpoints(endpoints)

	//wrap endpoints
	loggedEndpoints := Log(endpoints, log)

	http.Handle("/", loggedEndpoints)

	//start server
	log.Fatal(http.ListenAndServe(conf.Server.Address+":8080", nil))
}

// Log wrapper function for the http server
func Log(handler *mux.Router, log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
