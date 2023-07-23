package db

/*
	Test for:
		- Return the KeyHeader and not a random one
*/
func (r *Row) GetKeyHeaderAndValue() (HeaderI, ValueI) {
	for h, v := range r.RowMap {
		if h.IsKeyHeader() {
			return h, v
		}
	}

	return nil, nil
}

/*
	Test for:
		- Return the correct value for the given header
*/
func (r *Row) GetValueFromHeader(header string) ValueI {
	for h, v := range r.RowMap {
		if h.GetName() == header {
			return v
		}
	}

	return nil
}

/*
	Test for:
		- Return the correct value for the RowMap
*/
func (r *Row) GetRowMap() map[HeaderI]ValueI {
	return r.RowMap
}

/*
	Test for:
		- When the header exists
		- When the header doesn't exist
*/
func (r *Row) HeaderExists(header string) bool {
	for h := range r.RowMap {
		if h.GetName() == header {
			return true
		}
	}

	return false
}

/*
	Test for:
		- When the value is equal to the KeyHeader's
		- When the value isn't equal
*/
func (r *Row) KeyHeaderValueEqual(value string) bool {
	for h, v := range r.RowMap {
		if h.IsKeyHeader() {
			if v.GetValue() == value {
				return true
			}
		}
	}

	return false
}

/*
	Test for:
		- The value is added to the header correctly
		- If the header is marked as KeyHeader and one already exists
*/
func (r *Row) AddHeaderWithValue(header string, keyHeader bool, t Type, value string) {
	// TODO: update to check if a key header already exists
	h := &Header{header, keyHeader, t}
	v := &Value{value}
	r.RowMap[h] = v
}

/*
	Test for:
		- The header is removed correctly
*/
func (r *Row) RemoveHeader(header string) {
	// Ignores if header is Key Header
	newRowMap := map[HeaderI]ValueI{}
	for h, v := range r.RowMap {
		if h.GetName() != header {
			newRowMap[h] = v
		}
	}
}

/*
	Test for:
		- The header's value is updated correctly
*/
func (r *Row) UpdateHeaderValue(header string, value string) {
	for h := range r.RowMap {
		if h.GetName() == header {
			r.RowMap[h].SetValue(value)
			break
		}
	}
}
