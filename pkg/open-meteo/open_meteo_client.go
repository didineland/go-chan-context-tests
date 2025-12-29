package openmeteo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	url string = "https://api.open-meteo.com/v1/forecast?latitude=45.51770642278042&longitude=-73.56301971537141&current=temperature_2m,wind_speed_10m&hourly=temperature_2m,relative_humidity_2m,wind_speed_10m"
)

func GetMeteo(ctx context.Context, responseChannel chan<- Current) {
	for true {
		select {
		case <-ctx.Done():
			fmt.Println("Task completed or deadline exceeded:", ctx.Err())
			return
		default:
			callMeteo(ctx, responseChannel)
			time.Sleep(3 * time.Second)
		}
	}
}

func callMeteo(ctx context.Context, responseChannel chan<- Current) {
	resp := &WeatherResponse{}
	if apiCall(ctx, url, resp) {
		responseChannel <- resp.Current
	}
}

func apiCall(ctx context.Context, url string, response any) bool {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return false
	}
	client := http.DefaultClient
	resp, err := client.Do(req)

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		errDecode := json.NewDecoder(resp.Body).Decode(response)

		if errDecode != nil {
			return false
		}
	} else {

		errString := fmt.Sprintf("Error %d in API call", resp.StatusCode)
		err = errors.New(errString)

		log.Print(err)
	}

	return true
}
