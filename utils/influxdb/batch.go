package influxdb

import (
	"fmt"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

var wait time.Duration = time.Second

type BatchPoints struct {
	bps   map[string]chan *client.Point
	using bool
}

func (bps *BatchPoints) addConfigType(retentionPolicy string) {
	if bps.bps == nil {
		bps.bps = make(map[string]chan *client.Point)
	}
	bps.checkAndWait()
	bps.using = true
	bps.bps[retentionPolicy] = startPointStorage(&client.BatchPointsConfig{
		Database:        schema,
		Precision:       "s",
		RetentionPolicy: retentionPolicy,
	})
	bps.using = false
}

func startPointStorage(storage *client.BatchPointsConfig) chan *client.Point {
	return storePoints(storage, wait)
}

func (bps *BatchPoints) checkAndWait() {
	for bps.using {
		time.Sleep(time.Second)
	}
}

func storePoints(storage *client.BatchPointsConfig, storeInterval time.Duration) chan *client.Point {
	ch := make(chan *client.Point, 10000)
	last := time.Now()
	bps, _ := client.NewBatchPoints(*storage)
	go func() {
		for {
			var err error
			for !last.Add(storeInterval).After(time.Now()) { //gather points for 1 second
				select {
				case pt, ok := <-ch:
					if ok {
						bps.AddPoint(pt)
					} else {
						fmt.Printf("\n\nChannel was closed @ %+v\n\n", time.Now().UTC().String())
						log.Warn("Channel closed!")
					}
				default:
					log.Info("No stream ready, moving on.")
				}
			}

			if len(bps.Points()) > 0 {
				fmt.Println("\nstoring")
				conn.Write(bps)
			}

			bps, err = client.NewBatchPoints(*storage)
			if err != nil {
				log.Fatal("Error during storepoints function at creating new batchpoints", err)
				return
			}
			last = time.Now()
		}
	}()

	return ch
}

func AddPoint(measurement, retentionPolicy string, tags map[string]string, fields map[string]interface{}, t ...time.Time) error {
	var err error
	var pt *client.Point
	if t == nil {
		pt, err = client.NewPoint(measurement, tags, fields, time.Now())
	} else {
		pt, err = client.NewPoint(measurement, tags, fields, t[0])
	}
	fmt.Printf("Added new point %+v\n", pt)
	if err != nil {
		log.Fatal("Error adding datapoint: ", err)
		return err
	}

	rtp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:        schema,
		Precision:       "s",
		RetentionPolicy: retentionPolicy,
	})

	rtp.AddPoint(pt)
	conn.Write(rtp)
	fmt.Println("point Added")
	return nil
}
