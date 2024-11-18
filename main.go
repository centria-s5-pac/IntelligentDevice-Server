package main

import (
	"fmt"
	cmdapi "helios/cmd/api"
	"time"
)

func main() {

	fmt.Println("Starting the Rest API server...")
	cmdapi.Main()

	for {
		time.Sleep(1 * time.Second)
	}
}
