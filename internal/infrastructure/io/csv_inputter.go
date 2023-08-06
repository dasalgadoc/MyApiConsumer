package io

import (
	"encoding/csv"
	"myApiController/internal/domain"
	"os"
)

type csvInputter struct{}

func NewCsvInputter() domain.DataInputter {
	return &csvInputter{}
}

func (c *csvInputter) Invoke(location string) (domain.Table, error) {
	file, err := os.Open(location)
	if err != nil {
		return domain.Table{}, err
	}
	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		return domain.Table{}, err
	}

	rows, err := reader.ReadAll()
	if err != nil {
		return domain.Table{}, err
	}

	table := domain.Table{
		Headers: headers,
		Rows:    rows,
	}
	return table, nil
}

func (c *csvInputter) InputterExtension() string {
	return ".csv"
}
