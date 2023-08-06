package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"myApiController/internal/domain"
	"testing"
)

type csvInputterTestScenario struct {
	test        *testing.T
	element     string
	table       domain.Table
	err         error
	expectedErr error
	inputter    csvInputter
	generic     domain.DataInputter
}

func TestFailedOpenedFile(t *testing.T) {
	s := startCsvInputterTestScenario(t)
	s.givenAnFileLocation("no-existing.csv")
	s.andExpectedError(fmt.Errorf("open no-existing.csv: no such file or directory"))
	s.whenCsvInputterIsInvoked()
	s.thenThereIsAnError()
}

func TestEmptyInputFile(t *testing.T) {
	s := startCsvInputterTestScenario(t)
	s.givenAnFileLocation("../../test/empty_input.csv")
	s.andExpectedError(fmt.Errorf("EOF"))
	s.whenCsvInputterIsInvoked()
	s.thenThereIsAnError()
}

func TestOnlyHeadersInputFile(t *testing.T) {
	s := startCsvInputterTestScenario(t)
	s.givenAnFileLocation("../../test/only_headers_input.csv")
	s.whenCsvInputterIsInvoked()
	s.thenThereIsNoError()
}

func TestCsvInputter_InputterExtension(t *testing.T) {
	s := startCsvInputterTestScenario(t)
	s.thenAssertInputterExtension()
}

func TestCsvInputterConstructor(t *testing.T) {
	s := startCsvInputterTestScenario(t)
	s.whenCsvInputterIsCreated()
	s.thenGenericAssertionIsOk()
}

/*-- steps --*/
func startCsvInputterTestScenario(t *testing.T) *csvInputterTestScenario {
	t.Parallel()
	return &csvInputterTestScenario{
		test: t,
	}
}

func (c *csvInputterTestScenario) givenAnFileLocation(location string) {
	c.element = location
}

func (c *csvInputterTestScenario) andExpectedError(err error) {
	c.expectedErr = err
}

func (c *csvInputterTestScenario) whenCsvInputterIsInvoked() {
	c.table, c.err = c.inputter.Invoke(c.element)
}

func (c *csvInputterTestScenario) whenCsvInputterIsCreated() {
	c.generic = NewCsvInputter()
}

func (c *csvInputterTestScenario) thenAssertInputterExtension() {
	assert.Equal(c.test, ".csv", c.inputter.InputterExtension())
}

func (c *csvInputterTestScenario) thenThereIsNoError() {
	assert.NoError(c.test, c.err)
}

func (c *csvInputterTestScenario) thenThereIsAnError() {
	assert.Error(c.test, c.err)
	assert.Equal(c.test, c.expectedErr.Error(), c.err.Error())
}

func (c *csvInputterTestScenario) thenGenericAssertionIsOk() {
	assert.IsType(c.test, &csvInputter{}, c.generic)
}
