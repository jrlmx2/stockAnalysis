package stream

import (
	"fmt"
	"strings"
	"time"

	"github.com/jrlmx2/stockAnalysis/model"
	"github.com/jrlmx2/stockAnalysis/utils/logger"
	"github.com/jrlmx2/stockAnalysis/utils/term"
)

func ProcessStreams(log *logger.Logger, in chan interface{}) chan model.Unmarshalable {
	out := make(chan model.Unmarshalable, 10000)

	go func() {
		for {
			if term.WasTerminated() {
				log.Warn("Thread was terminated")
				log.Info("Stream monitor was terminated.")
				return
			}
			select {
			case stream, ok := <-in:
				if ok {
					go streamListener(stream.(Stream), out, log)
					fmt.Println("Found new stream")
				} else {
					fmt.Printf("\n\nChannel was closed @ %+v\n\n", time.Now().UTC().String())
					log.Warn("Channel closed!")
				}
			default:
				log.Info("No stream ready, moving on.")
			}
			time.Sleep(time.Second)
		}
	}()

	return out
}

func streamListener(reader Stream, out chan model.Unmarshalable, log *logger.Logger) {
	content := ""
	for {
		if term.WasTerminated() {
			log.Info("Stream parser was terminated.")
			return
		}

		line, err := reader.Connection().ReadString('>')
		if err != nil {
			log.Error("Error reading from stream: %s", err)
			//connection was closed, try again then kill this thread
			reader.Reopen()
			fmt.Printf("\n\nOpened Stream.\n\n")
			return
		}

		sline := string(line)

		if strings.Contains(sline, "/") && strings.Contains(sline, "status") {
			content = ""
			continue
		}

		if strings.Contains(sline, "/") && (strings.Contains(sline, "quote") || strings.Contains(sline, "trade")) {
			content += sline

			parsedContent, err := unmarshal(strings.Trim(strings.Trim(content, "\n"), " "))
			if err != nil {
				log.Warn("Unmarshalling string %s failed with %s", content, err)
			}
			out <- parsedContent
			content = ""
		} else {
			content += sline
		}
	}
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
