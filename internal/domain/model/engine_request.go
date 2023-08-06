package model

type EngineRequest struct {
	SiteId            string
	Payer             int64
	Collector         int64
	TransactionAmount float32
	TotalPaidAmount   float32
	OperationType     string
	PaymentMethod     string
	PaymentType       string
	ProductId         string
	ProcessingMode    string
	TransactionType   string
	SplitterId        string
	SplitterType      string
	SubType           string
	PayMarketPlaceId  string
}
