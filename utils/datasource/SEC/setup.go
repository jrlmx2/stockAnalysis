package sec

import (
	"github.com/jasonlvhit/gocron"
	"github.com/jrlmx2/stockAnalysis/utils/config"
)

var uri string

var key string

var limit = 5000

var scheduler *gocron.Scheduler

func Setup(conf config.API) {
	uri = conf.URL
	key = conf.Key

	//something about secheduling
	scheduler := gocron.NewScheduler()
	scheduler.Every(1).Day().At("23:59").Do(resetDayLimit)
	<-scheduler.Start()
}

func resetDayLimit() {
	limit = 5000
}
