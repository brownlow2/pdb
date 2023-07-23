package db

var (
	keyHeaderIncorrect          = "key header '%s' incorrect, expected '%s'"
	keyHeaderEmptyError         = "key header '%s' must not be empty"
	headerNotExistError         = "header '%s' does not exist"
	keyHeaderAlreadyExistsError = "key header already exists"
	deleteKeyHeaderError        = "cannot delete key header '%s'"
	keyHeaderValueExistsError   = "row with key header '%s' and value '%s' already exists"
)

// DB is the interface for any DB implementations
type DB interface {
	// Returns the name of the DB
	GetName() string

	// Returns the KeyHeader for the DB
	GetKeyHeader() string

	// Adds a header to the DB, also adding the header to each row with an empty value
	AddHeader(header HeaderI)

	// Removes a header from the DB, also removing the header from each row
	RemoveHeader(header string) error

	// Adds a row to the DB, adding in any missing headers. If there are extra headers in the row
	// being added, AddRow() returns an error
	AddRow(row RowI) error

	// Returns all rows in the DB
	GetRows() []RowI

	// Adds a new value to a given header based on KeyHeader == key
	// Returns an error if the header does not exist
	AddValueToHeader(value string, header string, key string) error

	// Returns a row based on KeyHeader == value
	GetRowFromKeyHeader(value string) RowI

	// Returns a list of the rows that have header == value
	// Returns an error if the header doesn't exist
	GetRowsFromHeaderAndValue(header string, value string) ([]RowI, error)
}

// The implementation for DB holding the following fields:
// Name: The name of the DB implementation
// KeyHeader: The KeyHeader for this DB implementation
// Headers: The headers created for this DB implementation
// Rows: The rows currently in the DB implementation
type DBImpl struct {
	Name      string
	KeyHeader string
	Headers   map[HeaderI]struct{}
	Rows      RowsI
}

// RowsI is the interface for the rows in a DB
type RowsI interface {
	// Returns the rows in the DB
	GetRows() []RowI

	// Adds a row to the DB
	// Returns an error if the new row has the same KeyHeader value as another row
	AddRow(row RowI) error

	// Deletes a row from the DB based on the rows KeyHeader value
	DeleteRow(row RowI)

	// Removes the given header from each of the rows
	RemoveHeader(header string)

	// Adds the value to the given header based on KeyHeader == key
	AddValueToRowWithKeyHeader(value string, header string, key string)

	// Returns the row with KeyHeader == keyHeaderValue
	// Returns nil if the keyHeaderValue doesn't match any rows
	GetRowFromKeyHeader(keyHeaderValue string) RowI

	// Returns a list of rows that satisfy header's value == value
	// Returns an empty list if none satisfy it
	GetRowsFromHeaderAndValue(header string, value string) ([]RowI, error)
}

// The implementation of Rows holding the following fields:
// Items: The list of Row instances
type Rows struct {
	Items []RowI
}

// RowI is the interface for a row in the Rows list
type RowI interface {
	// Returns the KeyHeader instance and its Value
	// Returns nil, nil if a KeyHeader doesn't exist
	GetKeyHeaderAndValue() (HeaderI, ValueI)

	// Returns the Value instance for a given header in the row
	// Returns error if the header doesn't exist
	GetValueFromHeader(header string) (ValueI, error)

	// Returns the row map of headers to their values
	GetRowMap() map[HeaderI]ValueI

	// Returns true if the header exists in the row, false otherwise
	HeaderExists(header string) bool

	// Returns true if the value given is equal to the KeyHeader's value for the row
	KeyHeaderValueEqual(value string) bool

	// Adds a new header with a given value to the row
	// Returns error if the new header is labelled as the KeyHeader and one already exists
	AddHeaderWithValue(header string, keyHeader bool, t Type, value string) error

	// Removes the given header from the row
	// Returns an error if trying to delete the KeyHeader
	RemoveHeader(header string) error

	// Updates a given header's value in the row
	UpdateHeaderValue(header string, value string)
}

// The implementation of a row holding the following fields:
// RowMap: The map of headers to their values
type Row struct {
	RowMap map[HeaderI]ValueI
}

// The type given to a header with the following values:
// VALUE_STRING: The header's value is a string
// VALUE_NUMBER: The header's value is a number and can be used in LessThan/MoreThan operations
type Type int

const (
	VALUE_STRING Type = iota
	VALUE_NUMBER
)

// HeaderI is the interface for a header in a Row or DB
type HeaderI interface {
	// Returns the name of the header
	GetName() string

	// Returns the type of the header
	GetType() Type

	// Returns true if the header's value is a string
	IsString() bool

	// Returns true if the header's value is a number
	IsNumber() bool

	// Returns true if the header has been assigned the KeyHeader role
	IsKeyHeader() bool
}

// The implementation of a header holding the following fields:
// Name: The name of the header
// KeyHeader: A bool denoted if the header is the key header in the row and DB
// Type: The type of the header
type Header struct {
	Name      string
	KeyHeader bool
	Type
}

// ValueI is the interface for a value in a Row or DB for a header
type ValueI interface {
	// Returns the value as a string
	GetValue() string

	// Sets the value for the intance
	SetValue(value string)
}

// The implementation of a value holding the following fields:
// Value: the value of the instance represented as a string
type Value struct {
	Value string
}
