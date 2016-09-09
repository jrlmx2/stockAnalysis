package streaming

import (
	"fmt"
	"strings"
	"time"

	"github.com/jrlmx2/stockAnalysis/model"
	"github.com/jrlmx2/stockAnalysis/utils/term"
	"github.com/op/go-logging"
)

const (
	open  = 7
	close = 17
)

func ProcessStreams(log *logging.Logger) <-chan model.Unmarshalable {
	out := make(chan model.Unmarshalable, 0)

	d, _ := time.ParseDuration("1s")
	go func() {
		for {
			if term.WasTerminated() {
				log.Warning("Thread was terminated")
				log.Info("Stream monitor was terminated.")
				return
			}
			select {
			case stream, ok := <-handler:
				if ok {
					go streamListener(stream, &out, log)
					fmt.Println("Found new stream")
				} else {
					fmt.Printf("\n\nChannel was closed @ %+v\n\n", time.Now().UTC().String())
					log.Warning("Channel closed!")
				}
			default:
				log.Info("No stream ready, moving on.")
			}
			time.Sleep(d)
		}
	}()

	return out
}

func streamListener(reader *TradeKingStream, out *chan model.Unmarshalable, log *logging.Logger) {
	content := ""
	for {
		if term.WasTerminated() {
			log.Info("Stream parser was terminated.")
			return
		}

		line, err := reader.S.ReadString('>')
		if err != nil {
			fmt.Printf("Error reading from stream: %s\n\n", err)
			log.Errorf("Error reading from stream: %s", err)
			//connection was closed, try again then kill this thread
			reinitiateStream(reader)
			return
		}

		sline := string(line)

		if strings.Contains(sline, "/") && strings.Contains(sline, "status") {
			content = ""
			continue
		}

		if strings.Contains(sline, "/") && (strings.Contains(sline, "quote") || strings.Contains(sline, "trade")) {
			content += sline

			_, err := unmarshal(strings.Trim(strings.Trim(content, "\n"), " "))
			if err != nil {
				log.Warningf("Unmarshalling string %s failed with %s", content, err)
			}
			//*out <- parsedContent
			content = ""
		} else {
			content += sline
		}
	}
}

func reinitiateStream(stream *TradeKingStream) {
	easternUnitedStates, _ := time.LoadLocation("America/New_York")
	now := time.Now().In(easternUnitedStates)
	if now.Hour() == 17 || (now.Hour() == 16 && now.Minute() > 50) { // market closes
		var wait time.Duration
		if now.Weekday().String() == "Friday" { //if its friday, wait the weekend otherwise, wait overnight
			wait, _ = time.ParseDuration((string)(48+close-open) + "h")
		} else {
			wait, _ = time.ParseDuration((string)(close-open) + "h")
		}
		fmt.Println("Stream waiting for " + wait.String())
		time.Sleep(wait)
	}

	fmt.Println("Opening new stream.")
	OpenStream(stream.Req)
}

func unmarshal(in string) (model.Unmarshalable, error) {
	if strings.Contains(in, "quote") {
		q, _ := model.NewQuoteU().Unmarshal(in)
		err := q.Save()
		fmt.Printf("\n Got Quote: %+v", q)
		return q, err
	}

	if strings.Contains(in, "trade") {
		trade, _ := model.NewTradeU().Unmarshal(in)
		err := trade.Save()
		fmt.Printf("\n Got Trade: %+v", trade)
		return trade, err
	}

	fmt.Printf("XML not identified %s", in)
	return model.NewQuoteU(), nil
}
