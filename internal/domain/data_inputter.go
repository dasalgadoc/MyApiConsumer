package domain

type Table struct {
	Headers []string
	Rows    [][]string
}

type DataInputter interface {
	Invoke(location string) (Table, error)
	InputterExtension() string
}
