package quandl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/jrlmx2/GoCountry"
	"github.com/jrlmx2/stockAnalysis/utils/config"
	"github.com/jrlmx2/stockAnalysis/utils/influxdb"
	"github.com/jrlmx2/stockAnalysis/utils/logger"
)

var scheduler *gocron.Scheduler

const responseDateFormat = "2006-01-02"

// APIURL is the string replace version of the quandl resource url
var apiURL string

// APIKEY is the personal key associated with a user account
var apiKey string

var isRunning bool

// Example api request for OPEC Crude Oil Price
// APIURL - https://www.quandl.com/api/v3/datasets/OPEC/ORB.json?api_key=pxPgy1zxxg_kW1vY_PeX
func getDate(date string) (time.Time, error) {
	return time.Parse(responseDateFormat, date)
}

func url(code *code, dataType string) string {
	if code.lastUpdated.IsZero() {
		return fmt.Sprintf(apiURL+"?api_key="+apiKey, code.Endpoint(), dataType)
	}

	return fmt.Sprintf(apiURL+"?api_key="+apiKey+"&start_date=%s&end_date=%s", code.Endpoint(), dataType, code.LastUpdated().Format(responseDateFormat), time.Now().UTC().Format(responseDateFormat))

}

// Init starts the scheduler to run the Gather function every day at 11pm
func Init(conf config.API, log *logger.Logger) {
	apiKey = conf.Key
	apiURL = conf.URL
	go Gather(log)
	scheduler := gocron.NewScheduler()
	scheduler.Every(1).Day().At("23:00").Do(Gather, log)
	<-scheduler.Start()
}

// Gather Gets all commodity information from quandl
func Gather(log *logger.Logger) {
	if isRunning {
		return
	}
	isRunning = true
	//find latest data point
	//loop through codes and create commodity objects and store in Market
	//debug := make(map[string]string)
	for _, code := range codes {
		meta := make(map[string]string)
		if code.ShouldGetNew() {

			resp, err := Get(code)
			if err != nil {
				log.Warn("Failed to retrieve code %s\n", code)
				continue
			}

			code.SetFrequency(resp.Dataset.Frequency)
			meta["frequency"] = resp.Dataset.Frequency

			if resp.Contains("index") {
				meta["index"] = "1"
			} else {
				meta["index"] = "0"
			}

			meta["source"] = strings.ToLower(strings.Split(code.Endpoint(), "/")[0])
			meta["endpoint"] = strings.ToLower(strings.Split(code.Endpoint(), "/")[1])

			countries := gocountry.Search(&gocountry.Options{Full: true, CodeTwo: true}, resp.Dataset.Name+" : "+resp.Dataset.Description)
			out := make([]string, len(countries))
			for pos, str := range countries {
				out[pos] = str.String()
			}
			meta["countries"] = strings.Join(out, ",")

			var class string
			for _, commodity := range commodities {
				if resp.Contains(commodity) {
					class = cleanString(commodity)
					break
				}
			}
			if class != "" {
				meta["commodity"] = class

				go storeData("quandl", class, "commodity", meta, resp.Dataset.ColumnNames, resp.Dataset.Data)
			} else {
				log.Warn("Failed to find a commodity type in endpoint %s.\nText: %s", code, strings.ToLower(resp.Dataset.Description)+":"+strings.ToLower(resp.Dataset.Name))
			}
			//codess := strings.Split(code.Endpoint(), "/")
			//if len(resp.Dataset.ColumnNames) > 0 {
			//	debug[codess[0]] = cleanString(strings.Join(resp.Dataset.ColumnNames, ", "))
			//}

			//fmt.Printf("Call to %s succeed %s and has columns %s", resp.Dataset.Name, code, strings.Join(resp.Dataset.ColumnNames, ", "))

		}
	}
	/*for key, value := range debug {
		fmt.Printf("%s = %s\n", key, value)
	}*/
	isRunning = false
}

func storeData(source, assetClass, storage string, identifiers map[string]string, columns []string, data [][]interface{}) {
	if len(columns) != len(data[0]) {
		panic(fmt.Sprintf("Column definition does not match data definition %d!=%d", len(columns), len(data[0])))
	}
	for _, entry := range data {
		if influxdb.IsRunning() {
			date, info := buildData(columns, entry)
			influxdb.AddPoint(source+"_"+assetClass, storage, identifiers, info, date)
		} else {
			buildData(columns, entry)
			//fmt.Printf("Would have added point %s_%s in storage %s\n", source, assetClass, storage)
		}
	}

}

func buildData(columns []string, entry []interface{}) (time.Time, map[string]interface{}) {
	data := make(map[string]interface{})
	var date time.Time
	for key, column := range columns {
		column = cleanString(column)
		if column == "date" {
			date, _ = time.Parse(responseDateFormat, entry[key].(string))
		} else {
			data[column] = entry[key]
		}
	}
	data["fetchTracker"] = 1
	return date, data
}

func cleanString(in string) string {
	in = strings.ToLower(in)
	in = strings.Replace(in, " ", "_", -1)
	in = strings.Replace(in, ",", "", -1)
	in = strings.Replace(in, ":", "", -1)
	in = strings.Replace(in, "-", "", -1)
	return strings.Replace(in, ".", "", -1)
}

// Response defines a quandl response object
type Response struct {
	Dataset struct {
		ID                  int             `json:"id"`
		DatasetCode         string          `json:"dataset_code"`
		DatabaseCode        string          `json:"database_code"`
		Name                string          `json:"name"`
		Description         string          `json:"description"`
		RefreshedAt         time.Time       `json:"refreshed_at"`
		NewestAvailableDate string          `json:"newest_available_date"`
		OldestAvailableDate string          `json:"oldest_available_date"`
		ColumnNames         []string        `json:"column_names"`
		Frequency           string          `json:"frequency"`
		Type                string          `json:"type"`
		Premium             bool            `json:"premium"`
		StartDate           string          `json:"start_date"`
		EndDate             string          `json:"end_date"`
		Data                [][]interface{} `json:"data"`
		DatabaseID          int             `json:"database_id"`
	} `json:"dataset"`
}

// Contains searchs selected pieces of the response object for string
func (r *Response) Contains(value string) bool {
	search := strings.ToLower(r.Dataset.Description) + " : " + strings.ToLower(r.Dataset.Name)
	if strings.Contains(search, " "+strings.ToLower(value)+".") || strings.Contains(search, " "+strings.ToLower(value)+" ") || strings.Contains(search, " "+strings.ToLower(value)+",") {
		return true
	}
	return false
}

// Get calls quandl with the url string code
func Get(code *code) (*Response, error) {
	//fmt.Println(url(code, "json"))
	resp, err := http.Get(url(code, "json"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	response := &Response{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
