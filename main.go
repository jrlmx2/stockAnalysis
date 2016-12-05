package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/hydrogen18/stoppableListener"
	logging "github.com/op/go-logging"

	"github.com/jrlmx2/stockAnalysis/API/tradeking"
	"github.com/jrlmx2/stockAnalysis/model"
	"github.com/jrlmx2/stockAnalysis/utils/config"
	"github.com/jrlmx2/stockAnalysis/utils/logger"
	"github.com/jrlmx2/stockAnalysis/utils/mariadb"
	"github.com/jrlmx2/stockAnalysis/utils/term"
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
	defer db.Close()

	model.SetRepository(db)

	log.Info("Database connected")

	//thread with a daily backup for the database
	go db.BackupDB(conf.Dump, conf.Database)

	//go model.MonitorWatchlists()

	// Setup Server
	listener, err := net.Listen("tcp", conf.Server.Address)
	if err != nil {
		term.Kill()
		panic(err)
	}

	sl, err := stoppableListener.New(listener)
	if err != nil {
		term.Kill()
		panic(err)
	}

	streamHandler := make(chan Stream)

	//establish endpoints
	endpoints := mux.NewRouter()
	endpoints = tradeking.EstablishEndpoints(endpoints)

	//wrap endpoints
	loggedEndpoints := Log(endpoints, log)

	http.Handle("/", loggedEndpoints)

	stop := term.Channel()
	server := &http.Server{
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 10 * time.Second,
	}

	// safely start server
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Serve(sl)
	}()

	fmt.Printf("Serving HTTP\n")
	select {
	case signal := <-stop:
		fmt.Printf("Got signal:%v\n", signal)
	}
	fmt.Printf("Stopping listener\n")
	sl.Stop()
	fmt.Printf("Waiting on server\n")
	wg.Wait()

}

// Log wrapper function for the http server
func Log(handler *mux.Router, log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		r.ParseForm()
		handler.ServeHTTP(w, r)
	})
}
