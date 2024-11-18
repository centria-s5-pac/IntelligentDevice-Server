package main

import (
	"fmt"
	cmdapi "helios/cmd/api"
)

func main() {
	fmt.Println("Starting the Rest API server...")
	go cmdapi.Main()

}
