package mocks

import (
	"github.com/stretchr/testify/mock"
	"myApiController/internal/domain"
)

type DataRowClientMock struct {
	mock.Mock
}

func (m *DataRowClientMock) DoRequest(params map[string]string, body string) (domain.DataExchange, error) {
	args := m.Called(params, body)
	return args.Get(0).(domain.DataExchange), args.Error(1)
}
