package configs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type configTestScenario struct {
	test         *testing.T
	fileLocation string
	config       Config
	err          error
	expectedErr  error
}

func TestLoadANoExistingConfig(t *testing.T) {
	s := startConfigTestScenario(t)
	s.givenAConfigFile("./xxx.yaml")
	s.andExpectedError(fmt.Errorf("open ./xxx.yaml: no such file or directory"))
	s.whenLoadingConfig()
	s.thenThereIsAnError()
}

func TestLoadAConfig(t *testing.T) {
	s := startConfigTestScenario(t)
	s.givenAConfigFile("./secure_config.yaml")
	s.whenLoadingConfig()
	s.thenThereIsNoError()
	s.thenAssertIO(IO{
		FolderLocation: "<YOUR FOLDER HERE>",
		InputFileName:  "<YOUR FILE HERE>",
	})
	s.thenAssertClients([]Client{
		{
			Name:    "api-engine",
			Path:    "<YOUR URL HERE>",
			Headers: map[string]string{"<HEADER KEY>": "<HEADER VALUE>"},
		},
		{
			Name: "no-existing-client",
		},
	})
}

func TestLoadInvalidConfigFile(t *testing.T) {
	s := startConfigTestScenario(t)
	s.givenAConfigFile("./invalid.yaml")
	s.andExpectedError(fmt.Errorf("yaml: unmarshal errors:\n  line 2: cannot unmarshal !!map into []configs.Client"))
	s.whenLoadingConfig()
	s.thenThereIsAnError()
}

/*-- steps --*/
func startConfigTestScenario(t *testing.T) *configTestScenario {
	t.Parallel()
	return &configTestScenario{
		test: t,
	}
}

func (c *configTestScenario) givenAConfigFile(location string) {
	c.fileLocation = location
}

func (c *configTestScenario) andExpectedError(err error) {
	c.expectedErr = err
}

func (c *configTestScenario) whenLoadingConfig() {
	c.config, c.err = LoadConfig(c.fileLocation)
}

func (c *configTestScenario) thenThereIsAnError() {
	assert.Error(c.test, c.err)
	assert.Equal(c.test, c.expectedErr.Error(), c.err.Error())
}

func (c *configTestScenario) thenThereIsNoError() {
	assert.NoError(c.test, c.err)
}

func (c *configTestScenario) thenAssertIO(expected IO) {
	assert.Equal(c.test, expected, c.config.IO)
}

func (c *configTestScenario) thenAssertClients(expected []Client) {
	assert.Equal(c.test, expected, c.config.Clients)
}
