package streaming

import (
	"bufio"
	"fmt"
	"strings"
	"time"

	"github.com/jrlmx2/stockAnalysis/model"
	"github.com/jrlmx2/stockAnalysis/utils/term"
	"github.com/op/go-logging"
)

func ProcessStreams(log *logging.Logger) <-chan model.Unmarshalable {
	out := make(chan model.Unmarshalable, 0)

	d, _ := time.ParseDuration("1s")
	go func() {
		for {
			if term.WasTerminated() {
				log.Info("Stream monitor was terminated.")
				return
			}
			select {
			case stream, ok := <-handler:
				if ok {
					go streamListener(stream, &out, log)
					fmt.Println("Found new stream")
				} else {
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

func streamListener(reader *bufio.Reader, out *chan model.Unmarshalable, log *logging.Logger) {
	content := ""
	for {
		if term.WasTerminated() {
			log.Info("Stream parser was terminated.")
			return
		}

		line, err := reader.ReadString('>')
		if err != nil {
			log.Errorf("Error reading from stream: %s", err)
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

func unmarshal(in string) (model.Unmarshalable, error) {
	fmt.Printf("\nGot in: %s\n\n", in)
	if strings.Contains(in, "quote") {
		q, _ := model.NewQuoteU().Unmarshal(in)
		err := q.Save()
		return q, err
	}

	if strings.Contains(in, "trade") {
		trade, _ := model.NewTradeU().Unmarshal(in)
		err := trade.Save()
		return trade, err
	}

	fmt.Printf("XML not identified %s", in)
	return model.NewQuoteU(), nil
}
