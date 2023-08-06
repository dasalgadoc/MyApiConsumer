package mocks

import (
	"github.com/stretchr/testify/mock"
	"myApiController/internal/domain"
)

type DataOutputterMock struct {
	mock.Mock
}

func (m *DataOutputterMock) Write(location string, rows []domain.DataExchange) error {
	args := m.Called(location, rows)
	return args.Error(0)
}

func (m *DataOutputterMock) OutputterFilename() string {
	args := m.Called()
	return args.Get(0).(string)
}
