package streaming

import (
	"fmt"
	"strings"
	"sync"

	"github.com/jrlmx2/stockAnalysis/model"
	"github.com/jrlmx2/stockAnalysis/utils/logger"
)

func ProcessStream(ch chan<- buffio.Reader, log logger.Logger) {
	var wg sync.WaitGroup

	go func() {
		for {
			select {
			case stream, ok := <-ch:
				if ok {
					go streamListener(stream, out, wg)
					wg.Add(1)
				} else {
					fmt.Println("Channel closed!")
				}
			default:
				fmt.Println("No value ready, moving on.")
			}
			wg.Done()
		}
	}()
	wg.Add(1)

	//send to listeners
	//send to database
}

func streamListener(reader buffio.Reader, out chan<- model.Unmarshalable, wg sync.WaitGroup, log logger.Logger) {
	content := ""
	for {
		line, err := reader.ReadString('>')
		if err != nil {
			return err
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
				log.Warn
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

	if strings.Contains(in, "quote") {
		return model.NewQuote().unmarshal(in)
	}

	if strings.Contains(in, "trade") {
		return model.NewTrade().unmarshal(in)
	}

	fmt.Printf("XML not identified %s", in)
	return model.NewQuote(), nil
}
