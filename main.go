package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	cmdapi "helios/cmd/api"
	"helios/common"
	"helios/lightbrain"
)

func main() {

	fmt.Println("Starting the Rest API server...")
	go cmdapi.Main()

	fmt.Println("Starting the MQTT broadcast instance...")
	go common.BroadcastServerIP()

	go lightbrain.Main()

	// * Wait for a signal to shutdown the server *
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	fmt.Println("Shutting down the Rest API server...")

}
