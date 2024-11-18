package main

import (
	"fmt"
	cmdapi "helios/cmd/api"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	fmt.Println("Starting the Rest API server...")
	go cmdapi.Main()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	fmt.Println("Shutting down the Rest API server...")

}
