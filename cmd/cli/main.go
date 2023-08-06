package main

import (
	"flag"
	"fmt"
	bootstrap "myApiController/cmd/bootstrap"
	"myApiController/cmd/configs"
)

// registered
var (
	IOsType = []string{configs.CsvIoType, configs.JsonIoType}
)

func main() {
	inputter := flag.String("inputter", "csv", "Inputter type")
	outputter := flag.String("outputter", "json", "Outputter type")
	flag.Parse()

	if !bootstrap.CheckArgumentOnSlice(*inputter, IOsType) {
		panic(fmt.Errorf("invalid inputter, the valid types are %+v", IOsType))
	}

	if !bootstrap.CheckArgumentOnSlice(*outputter, IOsType) {
		panic(fmt.Errorf("invalid outputter, the valid types are %+v", IOsType))
	}

	args := flag.Args()
	if len(args) < 1 {
		panic("no client provided")
	}
	client := &args[0]

	var app = bootstrap.BuildApplication(*inputter, *outputter, *client)

	registeredClientsNames := app.AppConfig.GetRegisteredClientsNames()
	if !bootstrap.CheckArgumentOnSlice(*client, registeredClientsNames) {
		panic(fmt.Errorf("invalid client, the valid types are %+v", registeredClientsNames))
	}

	fmt.Printf("...Running request from (%s), recovery data from (%s) and writing to (%s)\n",
		*client, *inputter, *outputter)
	app.DataProcessor.Do()
}
