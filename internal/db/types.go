package db

type DB interface {
	GetName() string
	AddHeader(header string)
	RemoveHeader(header string) error
	AddRow(row map[string]string) error
	GetRows() []map[string]string
}

type DBImpl struct {
	Name      string
	KeyHeader string
	Headers   map[string]struct{}
	Rows
}

type Rows struct {
	Items []RowI
}

type RowI interface {
	GetKeyHeaderAndValue() (HeaderI, ValueI)
	GetValueFromHeader(header string) ValueI
	HeaderExists(header string) bool
	KeyHeaderValueEqual(value string) bool
	AddHeaderWithValue(header string, keyHeader bool, t Type, value string)
	RemoveHeader(header string)
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
}

type Value struct {
	Value string
}
