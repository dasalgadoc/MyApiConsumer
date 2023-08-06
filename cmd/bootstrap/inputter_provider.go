package bootstrap

import (
	"fmt"
	"myApiController/cmd/configs"
	"myApiController/internal/domain"
	"myApiController/internal/infrastructure/io"
)

var inputter = map[string]func() (domain.DataInputter, error){
	configs.CsvIoType: buildCSVInputter,
}

func GetDataInputter(inputterType string) (domain.DataInputter, error) {
	i, exists := inputter[inputterType]
	if !exists {
		return nil, fmt.Errorf("unable to build %s outputter", inputterType)
	}

	return i()
}

func buildCSVInputter() (domain.DataInputter, error) {
	return io.NewCsvInputter(), nil
}
