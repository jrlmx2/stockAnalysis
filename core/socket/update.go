package socket

import (
	"fmt"
	"time"

	"github.com/jrlmx2/stockAnalysis/model"
	"github.com/jrlmx2/stockAnalysis/utils/term"
	logging "github.com/op/go-logging"
)

func UpdateSubscribers(log *logging.Logger, in chan model.Unmarshalable) {
	go func() {
		for {
			if term.WasTerminated() {
				log.Warning("Thread was terminated")
				log.Info("Stream monitor was terminated.")
				return
			}
			select {
			case _, ok := <-in:
				if ok {
					fmt.Println("Found new unmarshalable")
				} else {
					fmt.Printf("\n\nChannel was closed @ %+v\n\n", time.Now().UTC().String())
					log.Warning("Channel closed!")
				}
			default:
				log.Info("No stream ready, moving on.")
			}
		}
	}()
}
