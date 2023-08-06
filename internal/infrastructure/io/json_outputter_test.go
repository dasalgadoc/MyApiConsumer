package io

import (
	"github.com/stretchr/testify/assert"
	domain2 "myApiController/internal/domain"
	"testing"
)

type jsonOutputterTestScenario struct {
	test       *testing.T
	folderPath string
	filename   string
	rows       []domain2.DataExchange
	err        error
	outputter  jsonOutputter
	generic    domain2.DataOutputter
}

func TestCreateFileFailed(t *testing.T) {
	s := startJsonOutputterTestScenario(t)
	s.givenALocation("/dev/null/")
	s.andInputStruct([]domain2.DataExchange{"one"})
	s.whenJsonWriting()
	s.thenThereIsAnErrorNoValidateText()
}

func TestFileOutputOk(t *testing.T) {
	s := startJsonOutputterTestScenario(t)
	s.givenALocation("../../test/")
	s.andInputStruct([]domain2.DataExchange{"one"})
	s.whenJsonWriting()
	s.thenThereIsNoError()
}

func TestJsonOutputter_OutputterFilename(t *testing.T) {
	s := startJsonOutputterTestScenario(t)
	s.whenFilenameIsRequested()
	s.thenAssertOutputterFileName()
}

func TestJsonOutputterConstructor(t *testing.T) {
	s := startJsonOutputterTestScenario(t)
	s.whenJsonOutputterIsCreated()
	s.thenGenericAssertionIsOk()
}

/*-- steps --*/
func startJsonOutputterTestScenario(t *testing.T) *jsonOutputterTestScenario {
	t.Parallel()
	return &jsonOutputterTestScenario{
		test: t,
	}
}

func (j *jsonOutputterTestScenario) givenALocation(location string) {
	j.folderPath = location + j.outputter.OutputterFilename()
}

func (j *jsonOutputterTestScenario) andInputStruct(rows []domain2.DataExchange) {
	j.rows = rows
}

func (j *jsonOutputterTestScenario) whenFilenameIsRequested() {
	j.filename = j.outputter.OutputterFilename()
}

func (j *jsonOutputterTestScenario) whenJsonWriting() {
	j.err = j.outputter.Write(j.folderPath, j.rows)
}

func (j *jsonOutputterTestScenario) whenJsonOutputterIsCreated() {
	j.generic = NewJsonOutputter()
}

func (j *jsonOutputterTestScenario) thenAssertOutputterFileName() {
	prefix := j.filename[:7]
	extension := j.filename[len(j.filename)-5:]
	assert.Equal(j.test, "output_", prefix)
	assert.Equal(j.test, ".json", extension)
}

func (j *jsonOutputterTestScenario) thenGenericAssertionIsOk() {
	assert.IsType(j.test, &jsonOutputter{}, j.generic)
}

func (j *jsonOutputterTestScenario) thenThereIsNoError() {
	assert.NoError(j.test, j.err)
}

func (j *jsonOutputterTestScenario) thenThereIsAnErrorNoValidateText() {
	assert.Error(j.test, j.err)
}
