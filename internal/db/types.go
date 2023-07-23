package db

type DB interface {
	GetName() string
	GetKeyHeader() string
	AddHeader(header HeaderI)
	RemoveHeader(header string) error
	AddRow(row RowI) error
	GetRows() []RowI
	AddValueToHeader(value string, header string, key string) error
	GetRowFromKeyHeader(value string) RowI
	GetRowsFromHeaderAndValue(header string, value string) []RowI
}

type DBImpl struct {
	Name      string
	KeyHeader string
	Headers   map[HeaderI]struct{}
	Rows      RowsI
}

type RowsI interface {
	GetRows() []RowI
	AddRow(row RowI) error
	DeleteRow(row RowI)
	RemoveHeader(header string)
	AddValueToRowWithKeyHeader(value string, header string, key string)
	GetRowFromKeyHeader(keyHeaderValue string) RowI
	GetRowsFromHeaderAndValue(header string, value string) []RowI
}

type Rows struct {
	Items []RowI
}

type RowI interface {
	GetKeyHeaderAndValue() (HeaderI, ValueI)
	GetValueFromHeader(header string) ValueI
	GetRowMap() map[HeaderI]ValueI
	HeaderExists(header string) bool
	KeyHeaderValueEqual(value string) bool
	AddHeaderWithValue(header string, keyHeader bool, t Type, value string)
	RemoveHeader(header string)
	UpdateHeaderValue(header string, value string)
}

type Row struct {
	RowMap map[HeaderI]ValueI
}

type Type int

const (
	VALUE_STRING Type = iota
	VALUE_NUMBER
)

type HeaderI interface {
	GetName() string
	GetType() Type
	IsString() bool
	IsNumber() bool
	IsKeyHeader() bool
}

type Header struct {
	Name      string
	KeyHeader bool
	Type
}

type ValueI interface {
	GetValue() string
	SetValue(value string)
}

type Value struct {
	Value string
}
