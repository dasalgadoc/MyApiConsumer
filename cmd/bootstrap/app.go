package bootstrap

import (
	"fmt"
	"myApiController/cmd/configs"
	"myApiController/internal/application"
	domain2 "myApiController/internal/domain"
	"os"
)

type (
	Application struct {
		DataProcessor application.DataProcessor
		AppConfig     configs.Config
	}
)

func BuildApplication(inputterType, outputterType, clientType string) *Application {
	appConfig := getConfiguration()
	fmt.Printf("...Configuration loaded %+v\n", appConfig)

	inputter, err := buildInputter(inputterType)
	if err != nil {
		panic(fmt.Errorf("error building inputter: %w", err))
	}
	fmt.Println("...Inputter generated")

	outputter, err := buildOutputter(outputterType)
	if err != nil {
		panic(fmt.Errorf("error building outputter: %w", err))
	}
	fmt.Println("...Outputter generated")

	client, err := buildClients(clientType, appConfig)
	if err != nil {
		panic(fmt.Errorf("error building clients: %w", err))
	}
	fmt.Printf("...Client generated %+v\n", client)

	return &Application{
		DataProcessor: BuildDataProcessor(appConfig, inputter, outputter, client),
		AppConfig:     appConfig,
	}
}

func getConfiguration() configs.Config {
	fmt.Println(os.Getwd())
	appConfig, err := configs.LoadConfig("./cmd/configs/config.yaml")
	if err != nil {
		panic(fmt.Errorf("error getting configuration: %w", err))
	}
	return appConfig
}

func buildInputter(iType string) (domain2.DataInputter, error) {
	return GetDataInputter(iType)
}

func buildOutputter(oType string) (domain2.DataOutputter, error) {
	return GetDataOutputter(oType)
}

func buildClients(cType string, c configs.Config) (domain2.DataRowClient, error) {
	var client configs.Client
	for _, cli := range c.Clients {
		if cType == cli.Name {
			client = cli
		}
	}
	return GetDataRowClient(client)
}
