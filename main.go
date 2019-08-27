package main

import (
	"fmt"
	"github.com/antonholmquist/jason"
	"net/http"
)

func main() {
	resp, err := http.Get("http://weather.livedoor.com/forecast/webservice/json/v1?city=130010")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	execute(resp)
}

func execute(response *http.Response) {
	body, err := jason.NewObjectFromReader(response.Body)
	forecasts, err := body.GetObjectArray("forecasts")
	for _, forecast := range forecasts {
		date, _ := forecast.GetString("date")
		telop, _ := forecast.GetString("telop")
		fmt.Println(date)
		fmt.Println(telop)
	}
	if err != nil {
		panic(err)
	}
}
