package bootstrap

import (
	"myApiController/application"
	"myApiController/cmd/configs"
	"myApiController/domain"
)

func BuildDataProcessor(
	config configs.Config,
	inputter domain.DataInputter,
	outputter domain.DataOutputter,
	client domain.DataRowClient,
) application.DataProcessor {
	return application.NewDataProcessor(config, inputter, outputter, client)
}
