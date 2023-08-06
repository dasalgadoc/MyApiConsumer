package mocks

import (
	"github.com/stretchr/testify/mock"
	"myApiController/internal/domain"
)

type DataInputterMock struct {
	mock.Mock
}

func (m *DataInputterMock) Invoke(location string) (domain.Table, error) {
	args := m.Called(location)
	return args.Get(0).(domain.Table), args.Error(1)
}

func (m *DataInputterMock) InputterExtension() string {
	args := m.Called()
	return args.Get(0).(string)
}
