package db

/*
	Test for:
		- Correct name is returned
*/
func (h *Header) GetName() string {
	return h.Name
}

/*
	Test for:
		- Correct type is returned
*/
func (h *Header) GetType() Type {
	return h.Type
}

/*
	Test for:
		- If the header is a string, return true
		- If the header isn't, return false
*/
func (h *Header) IsString() bool {
	return h.Type == VALUE_STRING
}

/*
	Test for:
		- If the header is a number, return true
		- If the header isn't, return false
*/
func (h *Header) IsNumber() bool {
	return h.Type == VALUE_NUMBER
}

/*
	Test for:
		- When the header is the KeyHeader
		- When the header isn't the KeyHeader
*/
func (h *Header) IsKeyHeader() bool {
	return h.KeyHeader
}
