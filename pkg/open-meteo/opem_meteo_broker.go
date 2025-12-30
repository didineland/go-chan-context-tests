package openmeteo

import (
	"context"
	"fmt"

	chanbroker "github.com/didineland/meteo/pkg/chan-broker"
)

var (
	broker = chanbroker.NewBroker[Current]()
)

func StartMeteoBroker(ctx context.Context) {
	go broker.Start()
	meteoChan := make(chan Current)
	defer close(meteoChan)
	//ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	go GetMeteo(ctx, meteoChan)

	for {
		select {
		case current := <-meteoChan:
			fmt.Printf("Received: %+v\n", current)
			broker.Publish(current)
		case <-ctx.Done():
			fmt.Println("Broker cancelled")
			return
		}
	}
}

func RegisterListerner() chan Current {
	return broker.Subscribe()
}

func DeregisterListener(msgCh chan Current) {
	broker.Unsubscribe(msgCh)
}
