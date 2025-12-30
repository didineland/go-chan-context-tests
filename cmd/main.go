package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	openmeteo "github.com/didineland/meteo/pkg/open-meteo"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("Catched interrupt")
		cancel()
		os.Exit(1)
	}()

	go openmeteo.StartMeteoBroker(ctx)
	openmeteo.ConnectGrpcServer(ctx)

}
