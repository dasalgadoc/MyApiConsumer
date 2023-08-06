package domain

type DataRowClient interface {
	DoRequest(params map[string]string, body string) (DataExchange, error)
}
