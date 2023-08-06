package bootstrap

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type appTestScenario struct {
	test     *testing.T
	function func()
	output   *Application
}

func TestBuildApplicationFailed_Configs(t *testing.T) {
	s := startAppTestScenario(t)
	s.whenBuildApplicationExecuted("", "", "")
	s.thenTheAppplicationPanics()
}

func TestBuildApplicationOk(t *testing.T) {
	s := startAppTestScenario(t)
	err := os.Chdir("../../")
	if err != nil {
		panic(err)
	}
	s.whenBuildApplicationExecuted("csv", "json", "api-engine")
	s.thenTheApplicationReturnedOk()
}

func TestBuildApplicationFailed_Inputter(t *testing.T) {
	s := startAppTestScenario(t)
	s.whenBuildApplicationExecuted("", "", "")
	s.thenTheAppplicationPanics()
}

func TestBuildApplicationFailed_Outputter(t *testing.T) {
	s := startAppTestScenario(t)
	s.whenBuildApplicationExecuted("csv", "", "")
	s.thenTheAppplicationPanics()
}

func TestBuildApplicationFailed_Client(t *testing.T) {
	s := startAppTestScenario(t)
	s.whenBuildApplicationExecuted("csv", "json", "")
	s.thenTheAppplicationPanics()
}

/*-- steps --*/
func startAppTestScenario(t *testing.T) *appTestScenario {
	t.Parallel()
	return &appTestScenario{
		test: t,
	}
}

func (a *appTestScenario) whenBuildApplicationExecuted(inputter, outputter, client string) {
	a.function = func() {
		a.output = BuildApplication(inputter, outputter, client)
	}
}

func (a *appTestScenario) thenTheApplicationReturnedOk() {
	assert.NotPanics(a.test, a.function)
}

func (a *appTestScenario) thenTheAppplicationPanics() {
	assert.Panics(a.test, a.function)
}
