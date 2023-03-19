package providers

import (
	"fmt"
	"myApiController/configs"
	"myApiController/domain"
	"myApiController/infrastructure/io"
)

var outputter = map[string]func() (domain.DataOutputter, error){
	configs.JsonIoType: buildJsonOutputter,
}

func GetDataOutputter(outputterType string) (domain.DataOutputter, error) {
	o, exists := outputter[outputterType]
	if !exists {
		return nil, fmt.Errorf("unable to build %s outputter", outputterType)
	}

	return o()
}

func buildJsonOutputter() (domain.DataOutputter, error) {
	return io.NewJsonOutputter(), nil
}
