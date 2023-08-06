package bootstrap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type checkerTestScenario struct {
	test    *testing.T
	element string
	args    []string
	result  bool
}

func TestExistingElement(t *testing.T) {
	s := startCheckerTestScenario(t)
	s.givenAnElement("ABC")
	s.andSomeArgs([]string{"123", "ABC", "XXX"})
	s.whenCheckArgumentOnSlice()
	s.thenTrueResult()
}

func TestNoExistingElement(t *testing.T) {
	s := startCheckerTestScenario(t)
	s.givenAnElement("ABC")
	s.andSomeArgs([]string{"123", "XXX"})
	s.whenCheckArgumentOnSlice()
	s.thenFalseResult()
}

/*-- steps --*/
func startCheckerTestScenario(t *testing.T) *checkerTestScenario {
	t.Parallel()
	return &checkerTestScenario{
		test: t,
	}
}

func (c *checkerTestScenario) givenAnElement(element string) {
	c.element = element
}

func (c *checkerTestScenario) andSomeArgs(args []string) {
	c.args = args
}

func (c *checkerTestScenario) whenCheckArgumentOnSlice() {
	c.result = CheckArgumentOnSlice(c.element, c.args)
}

func (c *checkerTestScenario) thenFalseResult() {
	assert.False(c.test, c.result)
}

func (c *checkerTestScenario) thenTrueResult() {
	assert.True(c.test, c.result)
}
