package streaming

import (
	"bufio"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/jrlmx2/stockAnalysis/model"
	"github.com/jrlmx2/stockAnalysis/utils/term"
	"github.com/op/go-logging"
)

func ProcessStreams(log *logging.Logger) <-chan model.Unmarshalable {
	var wg sync.WaitGroup
	out := make(chan model.Unmarshalable, 0)

	die := term.NewTerm() //watch for sigterm, kill and others die will be 1 when terminated

	d, _ := time.ParseDuration("1s")
	go func() {
		for {
			if *die == 1 {
				log.Info("Stream monitor was terminated.")
				return
			}
			select {
			case stream, ok := <-handler:
				if ok {
					go streamListener(stream, out, &wg, die, log)
					wg.Add(1)
				} else {
					log.Warning("Channel closed!")
				}
			default:
				log.Info("No stream ready, moving on.")
			}
			wg.Done()
			time.Sleep(d)
		}
	}()
	wg.Add(1)

	return out
}

func streamListener(reader *bufio.Reader, out chan model.Unmarshalable, wg *sync.WaitGroup, die *int, log *logging.Logger) {
	content := ""
	for {
		if *die == 1 {
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
			fmt.Printf("CONTENT: %+v\n\n", content)
			parsedContent, err := unmarshal(content)
			if err != nil {
				log.Warningf("Unmarshalling string %s failed with %s", content, err)
			}
			out <- parsedContent
			content = ""
		} else {
			content += sline
		}
		wg.Done()
	}
}

func unmarshal(in string) (model.Unmarshalable, error) {
	fmt.Printf("\nGot in: %s\n\n", in)
	if strings.Contains(in, "quote") {
		return model.NewQuoteU().Unmarshal(in)
	}

	if strings.Contains(in, "trade") {
		trade, err := model.NewTradeU().Unmarshal(in)
		trade.Save()
		return trade, err
	}

	fmt.Printf("XML not identified %s", in)
	return model.NewQuoteU(), nil
}
