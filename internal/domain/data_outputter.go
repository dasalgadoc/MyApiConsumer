package domain

type DataOutputter interface {
	Write(string, []DataExchange) error
	OutputterFilename() string
}
