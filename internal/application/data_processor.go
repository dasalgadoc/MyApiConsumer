package application

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"myApiController/cmd/configs"
	domain2 "myApiController/internal/domain"
)

type DataProcessor struct {
	config    configs.Config
	inputter  domain2.DataInputter
	outputter domain2.DataOutputter
	rowClient domain2.DataRowClient
}

func NewDataProcessor(c configs.Config,
	i domain2.DataInputter,
	o domain2.DataOutputter,
	r domain2.DataRowClient) DataProcessor {
	return DataProcessor{
		config:    c,
		inputter:  i,
		outputter: o,
		rowClient: r,
	}
}

func (dp *DataProcessor) Do() {
	filepath := dp.config.IO.FolderLocation
	data, err := dp.inputter.Invoke(filepath + dp.config.IO.InputFileName + dp.inputter.InputterExtension())
	if err != nil {
		panic(fmt.Errorf("data inputter error: %w", err))
	}
	fmt.Printf("...Data read successfully from source. Has (%d) row(s)\n", len(data.Rows))

	dataReturned := dp.getDataFromRegisteredClient(data)

	err = dp.outputter.Write(filepath+dp.outputter.OutputterFilename(), dataReturned)
	if err != nil {
		panic(fmt.Errorf("data outputter error: %w", err))
	}
	fmt.Println("...Data was wrote successfully")
}

func (dp *DataProcessor) getDataFromRegisteredClient(data domain2.Table) []domain2.DataExchange {
	dataReturned := []domain2.DataExchange{}
	progressBar := progressbar.Default(int64(len(data.Rows)))
	for _, row := range data.Rows {
		var jsonBody string
		if data.Headers[0] == "JSON_BODY" {
			jsonBody = row[0]
		}
		params := dp.rowToParams(data.Headers, row)
		rowProcessed, err := dp.rowClient.DoRequest(params, jsonBody)
		if err != nil {
			continue
		}
		dataReturned = append(dataReturned, rowProcessed)
		progressBar.Add(1)
	}
	fmt.Printf("...Data recovery successfully from client. Has (%d) row(s)\n", len(dataReturned))
	return dataReturned
}

func (dp *DataProcessor) rowToParams(headers []string, row []string) map[string]string {
	params := make(map[string]string)
	for i, cell := range row {
		if headers[i] == "JSON_BODY" {
			continue
		}
		params[headers[i]] = cell
	}
	return params
}
