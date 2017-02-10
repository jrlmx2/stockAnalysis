package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/hydrogen18/stoppableListener"

	"github.com/jrlmx2/stockAnalysis/API/tradeking"
	"github.com/jrlmx2/stockAnalysis/core/socket"
	"github.com/jrlmx2/stockAnalysis/core/stream"
	"github.com/jrlmx2/stockAnalysis/model"
	"github.com/jrlmx2/stockAnalysis/utils/config"
	"github.com/jrlmx2/stockAnalysis/utils/influxdb"
	"github.com/jrlmx2/stockAnalysis/utils/logger"
	"github.com/jrlmx2/stockAnalysis/utils/mariadb"
	"github.com/jrlmx2/stockAnalysis/utils/quandl"
	"github.com/jrlmx2/stockAnalysis/utils/term"
)

func main() {

	//read config
	conf, file := config.ReadConfig()
	if conf == nil {
		panic("Config -c " + file + " could not be read.")
	}

	//setup logging
	log, err := logger.NewLogger("main", conf.Logger)
	if err != nil {
		panic(fmt.Sprintf("Logger failed to open with error:%s \n Configuration Details: %+v\n", err, conf.Logger))
	}

	//connect database
	db, err := mariadb.NewPool(conf.Database)
	if err != nil {
		log.Critical("Error opening mariadb at host %s", conf.Database.Host)
	}
	defer db.Close()

	c, err := influxdb.Setup(conf.InfluxDatabase, conf.Logger)
	if err != nil {
		log.Critical("Error opening influx at host %s", conf.InfluxDatabase.Host)
	}
	defer c.Close()

	model.SetRepository(db)
	modelLog, _ := logger.NewLogger("model", conf.Logger)
	model.SetLogger(modelLog)

	log.Info("Database connected")

	quandlLogger, err := logger.NewLogger("quandl", conf.Logger)
	if err != nil {
		panic(err)
	}
	go quandl.Init(conf.API["Quandl"], quandlLogger)

	//thread with a daily backup for the database
	//go db.BackupDB(conf.Dump, conf.Database)

	// Setup Server
	listener, err := net.Listen("tcp", conf.Server.Address)
	if err != nil {
		term.Kill()
		log.Fatal("Failed to create tcp listener with: ", err)
	}

	sl, err := stoppableListener.New(listener)
	if err != nil {
		term.Kill()
		log.Fatal("Failed to create tcp listener with: ", err)
	}

	streamHandler := make(chan interface{})
	subscribersLog, err := logger.NewLogger("FrontEndUpdates", conf.Logger)
	monitorLog, err := logger.NewLogger("TradekingMonitor", conf.Logger)
	socket.UpdateSubscribers(subscribersLog, stream.ProcessStreams(monitorLog, streamHandler))

	//establish endpoints
	endpoints := mux.NewRouter()
	endpoints = tradeking.EstablishEndpoints(endpoints, streamHandler)

	//wrap endpoints
	loggedEndpoints := Log(endpoints, log)

	http.Handle("/", loggedEndpoints)

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

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	fmt.Printf("Serving HTTP\n")
	select {
	case signal := <-sigc:
		fmt.Printf("Got signal:%v\n", signal)
	}
	fmt.Printf("Stopping listener\n")
	sl.Stop()
	fmt.Printf("Waiting on server\n")
	wg.Wait()

}

// Log wrapper function for the http server
func Log(handler *mux.Router, log *logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		r.ParseForm()
		handler.ServeHTTP(w, r)
	})
}
