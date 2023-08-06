package io

import (
	"encoding/json"
	domain2 "myApiController/internal/domain"
	"os"
	"time"
)

type jsonOutputter struct{}

func NewJsonOutputter() domain2.DataOutputter {
	return &jsonOutputter{}
}

func (c *jsonOutputter) Write(location string, rows []domain2.DataExchange) error {
	file, err := os.Create(location)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	for _, row := range rows {
		err = encoder.Encode(row)
		if err != nil {
			continue
		}
	}
	return nil
}

func (c *jsonOutputter) OutputterFilename() string {
	return "output_" + time.Now().String() + ".json"
}
