package service

import (
	"github.com/stretchr/testify/assert"
	"myApiController/internal/domain/model"
	"testing"
)

type engineRequestBuilderTestScenario struct {
	test     *testing.T
	request  model.EngineRequest
	expected model.EngineRequest
}

func TestFullBuilderObjet(t *testing.T) {
	s := startEngineRequestBuilderTestScenario(t)
	s.expected = model.EngineRequest{
		SiteId:            "SiteId",
		Payer:             123,
		Collector:         456,
		TransactionAmount: 1000,
		TotalPaidAmount:   100,
		OperationType:     "OP",
		PaymentMethod:     "PM",
		PaymentType:       "PT",
		ProductId:         "12345",
		ProcessingMode:    "ProcessingM",
		TransactionType:   "TT",
		SplitterId:        "67890",
		SplitterType:      "ST",
		SubType:           "SubT",
		PayMarketPlaceId:  "PayPI",
	}
	s.request = NewEngineRequestBuilder().
		SiteId("SiteId").
		Payer(123).
		Collector(456).
		TransactionAmount(1000).
		TotalPaidAmount(100).
		OperationType("OP").
		PaymentMethod("PM").
		PaymentType("PT").
		ProductId("12345").
		ProcessingMode("ProcessingM").
		TransactionType("TT").
		SplitterId("67890").
		SplitterType("ST").
		SubType("SubT").
		PayMarketPlaceId("PayPI").
		Build()
	s.thenObjectsMatch()
}

func TestEmptyBuilderObject(t *testing.T) {
	s := startEngineRequestBuilderTestScenario(t)
	s.expected = model.EngineRequest{}
	s.request = NewEngineRequestBuilder().Build()
	s.thenObjectsMatch()
}

func TestPartialBuilderObject(t *testing.T) {
	s := startEngineRequestBuilderTestScenario(t)
	s.expected = model.EngineRequest{
		SiteId:            "SiteId",
		Payer:             123,
		Collector:         456,
		TransactionAmount: 1500,
		OperationType:     "OP",
		PaymentMethod:     "PM",
		ProductId:         "123456",
	}
	s.request = NewEngineRequestBuilder().
		SiteId("SiteId").
		Payer(123).
		Collector(456).
		TransactionAmount(1500).
		OperationType("OP").
		PaymentMethod("PM").
		ProductId("123456").
		Build()
	s.thenObjectsMatch()
}

/*-- steps --*/
func startEngineRequestBuilderTestScenario(t *testing.T) *engineRequestBuilderTestScenario {
	t.Parallel()
	return &engineRequestBuilderTestScenario{
		test: t,
	}
}

func (e *engineRequestBuilderTestScenario) thenObjectsMatch() {
	assert.Equal(e.test, e.expected, e.request)
}
