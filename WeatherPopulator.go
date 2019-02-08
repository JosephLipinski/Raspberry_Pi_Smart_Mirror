package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Feature struct {
	Forecast10Day int
}

type Response struct {
	Version        string
	TermsOfService string
	Features       Feature
}
type ForecastDay []struct {
	Period         int
	Icon           string
	Icon_url       string
	Title          string
	Fcttext        string
	Fcttext_metric string
	Pop            string
}
type TxtForecast struct {
	Date        string
	Forecastday ForecastDay
}

type Date struct {
	Epoch           string
	Pretty          string
	Day             int
	Month           int
	Year            int
	Yday            int
	Hour            int
	Min             string
	Sec             int
	Isdst           int
	Monthname       string
	Monthname_short string
	Weekday_short   string
	Weekday         string
	Ampm            string
	Tz_short        string
	Tz_long         string
}
type High struct {
	Fahrenheit string
	Celsius    string
}

type Low struct {
	Fahrenheit string
	Celsius    string
}

type Qpf_allday struct {
	In float32
	Mm float32
}

type Qpf_day struct {
	In float32
	Mm float32
}

type Qpf_night struct {
	In float32
	Mm float32
}

type Snow_allday struct {
	In float32
	Cm float32
}

type Snow_day struct {
	In float32
	Cm float32
}

type Snow_night struct {
	In float32
	Cm float32
}

type Maxwind struct {
	Mph     int
	Kph     int
	Dir     string
	Degrees int
}
type Avewind struct {
	Mph     int
	Kph     int
	Dir     string
	Degrees int
}
type ForecastDay2 []struct {
	Date        Date
	Period      int
	High        High
	Low         Low
	Conditions  string
	Icon        string
	Icon_url    string
	Skyicon     string
	Pop         int
	Qpf_allday  Qpf_allday
	Qpf_day     Qpf_day
	Qpf_night   Qpf_night
	Snow_allday Snow_allday
	Snow_day    Snow_day
	Snow_night  Snow_night
	Maxwind     Maxwind
	Avewind     Avewind
	Avehumidity int
	Maxhumidity int
	Minhumidity int
}
type Simpleforecast struct {
	Forecastday ForecastDay2
}
type Forecast struct {
	Txt_forecast   TxtForecast
	Simpleforecast Simpleforecast
}

type Body struct {
	Response Response
	Forecast Forecast
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	url := "http://api.wunderground.com/api/1e404159ad02dddd/forecast10day/q/PA/Philadelphia.json"
	JSONResponse := Body{}
	getJson(url, &JSONResponse)
	fmt.Println(JSONResponse)
}
