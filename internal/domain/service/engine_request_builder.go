package service

import (
	"myApiController/internal/domain/model"
)

type EngineRequestBuilder struct {
	request model.EngineRequest
}

func NewEngineRequestBuilder() *EngineRequestBuilder {
	return &EngineRequestBuilder{request: model.EngineRequest{}}
}

func (rb *EngineRequestBuilder) SiteId(site string) *EngineRequestBuilder {
	rb.request.SiteId = site
	return rb
}

func (rb *EngineRequestBuilder) Payer(payer int64) *EngineRequestBuilder {
	rb.request.Payer = payer
	return rb
}

func (rb *EngineRequestBuilder) Collector(collector int64) *EngineRequestBuilder {
	rb.request.Collector = collector
	return rb
}

func (rb *EngineRequestBuilder) TransactionAmount(amount float32) *EngineRequestBuilder {
	rb.request.TransactionAmount = amount
	return rb
}

func (rb *EngineRequestBuilder) TotalPaidAmount(paidAmount float32) *EngineRequestBuilder {
	rb.request.TotalPaidAmount = paidAmount
	return rb
}

func (rb *EngineRequestBuilder) OperationType(operation string) *EngineRequestBuilder {
	rb.request.OperationType = operation
	return rb
}

func (rb *EngineRequestBuilder) PaymentMethod(paymentMethod string) *EngineRequestBuilder {
	rb.request.PaymentMethod = paymentMethod
	return rb
}

func (rb *EngineRequestBuilder) PaymentType(paymentType string) *EngineRequestBuilder {
	rb.request.PaymentType = paymentType
	return rb
}

func (rb *EngineRequestBuilder) ProductId(product string) *EngineRequestBuilder {
	rb.request.ProductId = product
	return rb
}

func (rb *EngineRequestBuilder) ProcessingMode(processingMode string) *EngineRequestBuilder {
	rb.request.ProcessingMode = processingMode
	return rb
}

func (rb *EngineRequestBuilder) TransactionType(transactionType string) *EngineRequestBuilder {
	rb.request.TransactionType = transactionType
	return rb
}

func (rb *EngineRequestBuilder) SplitterId(splitter string) *EngineRequestBuilder {
	rb.request.SplitterId = splitter
	return rb
}

func (rb *EngineRequestBuilder) SplitterType(splitterType string) *EngineRequestBuilder {
	rb.request.SplitterType = splitterType
	return rb
}

func (rb *EngineRequestBuilder) SubType(subType string) *EngineRequestBuilder {
	rb.request.SubType = subType
	return rb
}

func (rb *EngineRequestBuilder) PayMarketPlaceId(payMarketPlace string) *EngineRequestBuilder {
	rb.request.PayMarketPlaceId = payMarketPlace
	return rb
}

func (rb *EngineRequestBuilder) Build() model.EngineRequest {
	return rb.request
}
