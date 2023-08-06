package application

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"myApiController/cmd/configs"
	"myApiController/internal/domain"
	"myApiController/test/mocks"
	"testing"
)

var (
	headersWithoutJson = []string{"one", "two", "three"}
	rowsWithoutJson    = [][]string{{"1", "2", "3"}}

	headersWithJson = []string{"JSON_BODY", "one", "two", "three"}
	rowsWithJson    = [][]string{{`{"hello":"world"}`, "1", "2", "3"}}
)

type dataProcessorTestScenario struct {
	test          *testing.T
	function      func()
	configMock    configs.Config
	inputterMock  mocks.DataInputterMock
	outputterMock mocks.DataOutputterMock
	rowClientMock mocks.DataRowClientMock
}

func TestDataProcessorOK(t *testing.T) {
	s := startDataProcessorTestScenario(t)
	s.givenAConfig()
	s.andInputterIsOk(headersWithoutJson, rowsWithoutJson)
	s.andDataRowClientIsOk()
	s.andOutputterIsOk()
	s.whenDoingDataProcessor()
	s.thenThereIsNoPanics()
}

func TestDataProcessorWithBodyOk(t *testing.T) {
	s := startDataProcessorTestScenario(t)
	s.givenAConfigWithJsonBody()
	s.andInputterIsOk(headersWithJson, rowsWithJson)
	s.andDataRowClientIsOk()
	s.andOutputterIsOk()
	s.whenDoingDataProcessor()
	s.thenThereIsNoPanics()
}

func TestDataProcessorInputterFails(t *testing.T) {
	s := startDataProcessorTestScenario(t)
	s.givenAConfig()
	s.andInputterFailed()
	s.whenDoingDataProcessor()
	s.thenThereIsPanics()
}

func TestDataProcessorClientFails(t *testing.T) {
	s := startDataProcessorTestScenario(t)
	s.givenAConfig()
	s.andInputterIsOk(headersWithoutJson, rowsWithoutJson)
	s.andDataRowClientFailed()
	s.andOutputterIsOk()
	s.whenDoingDataProcessor()
	s.thenThereIsNoPanics()
}

func TestDataProcessorOutputterFails(t *testing.T) {
	s := startDataProcessorTestScenario(t)
	s.givenAConfig()
	s.andInputterIsOk(headersWithoutJson, rowsWithoutJson)
	s.andDataRowClientIsOk()
	s.andOutputterFailed()
	s.whenDoingDataProcessor()
	s.thenThereIsPanics()
}

/*--steps--*/
func startDataProcessorTestScenario(t *testing.T) *dataProcessorTestScenario {
	t.Parallel()
	return &dataProcessorTestScenario{
		test: t,
	}
}

func (d *dataProcessorTestScenario) givenAConfig() {
	d.configMock = configs.Config{
		IO: configs.IO{
			FolderLocation: "../test/",
			InputFileName:  "only_headers_input",
		},
	}
}

func (d *dataProcessorTestScenario) givenAConfigWithJsonBody() {
	d.configMock = configs.Config{
		IO: configs.IO{
			FolderLocation: "../test/",
			InputFileName:  "headers_with_json_body",
		},
	}
}

func (d *dataProcessorTestScenario) andInputterIsOk(headers []string, rows [][]string) {
	d.inputterMock.On("InputterExtension").Return(".csv")
	d.inputterMock.On("Invoke", mock.Anything).
		Return(
			domain.Table{
				Headers: headers,
				Rows:    rows,
			}, nil)
}

func (d *dataProcessorTestScenario) andInputterFailed() {
	d.inputterMock.On("InputterExtension").Return(".csv")
	d.inputterMock.On("Invoke", mock.Anything).
		Return(domain.Table{}, fmt.Errorf("something went wrong"))
}

func (d *dataProcessorTestScenario) andDataRowClientIsOk() {
	d.rowClientMock.On("DoRequest", mock.Anything, mock.Anything).
		Return("something", nil)
}

func (d *dataProcessorTestScenario) andDataRowClientFailed() {
	d.rowClientMock.On("DoRequest", mock.Anything, mock.Anything).
		Return("", fmt.Errorf("something went wrong"))
}

func (d *dataProcessorTestScenario) andOutputterIsOk() {
	d.outputterMock.On("OutputterFilename").Return("myFile.file")
	d.outputterMock.On("Write", mock.Anything, mock.Anything).
		Return(nil)
}

func (d *dataProcessorTestScenario) andOutputterFailed() {
	d.outputterMock.On("OutputterFilename").Return("myFile.file")
	d.outputterMock.On("Write", mock.Anything, mock.Anything).
		Return(fmt.Errorf("something went wrong"))
}

func (d *dataProcessorTestScenario) whenDoingDataProcessor() {
	target := NewDataProcessor(d.configMock, &d.inputterMock, &d.outputterMock, &d.rowClientMock)
	d.function = func() {
		target.Do()
	}
}

func (d *dataProcessorTestScenario) thenThereIsNoPanics() {
	assert.NotPanics(d.test, d.function)
}

func (d *dataProcessorTestScenario) thenThereIsPanics() {
	assert.Panics(d.test, d.function)
}
