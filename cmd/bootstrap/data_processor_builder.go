package bootstrap

import (
	"myApiController/cmd/configs"
	"myApiController/internal/application"
	domain2 "myApiController/internal/domain"
)

func BuildDataProcessor(
	config configs.Config,
	inputter domain2.DataInputter,
	outputter domain2.DataOutputter,
	client domain2.DataRowClient,
) application.DataProcessor {
	return application.NewDataProcessor(config, inputter, outputter, client)
}
