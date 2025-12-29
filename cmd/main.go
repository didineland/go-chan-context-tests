package main

import (
	"context"
	"fmt"
	"time"

	openmeteo "github.com/didineland/meteo/pkg/open-meteo"
)

func main() {

	meteoChan := make(chan openmeteo.Current)
	defer close(meteoChan)
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	go openmeteo.GetMeteo(ctx, meteoChan)

	for {
		select {
		case current := <-meteoChan:
			fmt.Printf("Received: %+v\n", current)
		case <-ctx.Done():
			return
		}
	}

}
