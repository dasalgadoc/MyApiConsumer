package main

import (
	"flag"
	"fmt"
	"myApiController/cmd/application"
	"myApiController/configs"
)

// registered
var (
	IOsType = []string{configs.CsvIoType, configs.JsonIoType}
)

func main() {
	inputter := flag.String("inputter", "csv", "Inputter type")
	outputter := flag.String("outputter", "json", "Outputter type")
	flag.Parse()

	if !application.CheckArgumentOnSlice(*inputter, IOsType) {
		panic(fmt.Errorf("invalid inputter, the valid types are %+v", IOsType))
	}

	if !application.CheckArgumentOnSlice(*outputter, IOsType) {
		panic(fmt.Errorf("invalid outputter, the valid types are %+v", IOsType))
	}

	args := flag.Args()
	if len(args) < 1 {
		panic("no client provided")
	}
	client := &args[0]

	var app = application.BuildApplication(*inputter, *outputter, *client)

	registeredClientsNames := app.AppConfig.GetRegisteredClientsNames()
	if !application.CheckArgumentOnSlice(*client, registeredClientsNames) {
		panic(fmt.Errorf("invalid client, the valid types are %+v", registeredClientsNames))
	}

	fmt.Printf("...Running request from (%s), recovery data from (%s) and writing to (%s)\n",
		*client, *inputter, *outputter)
	app.DataProcessor.Do()
}
