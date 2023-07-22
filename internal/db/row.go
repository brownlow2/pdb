package db

func (r *Row) GetKeyHeaderAndValue() (HeaderI, ValueI) {
	for h, v := range r.RowMap {
		if h.IsKeyHeader() {
			return h, v
		}
	}

	return nil, nil
}

func (r *Row) GetValueFromHeader(header string) ValueI {
	for h, v := range r.RowMap {
		if h.GetName() == header {
			return v
		}
	}

	return nil
}

func (r *Row) GetRowMap() map[HeaderI]ValueI {
	return r.RowMap
}

func (r *Row) HeaderExists(header string) bool {
	for h := range r.RowMap {
		if h.GetName() == header {
			return true
		}
	}

	return false
}

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

func (r *Row) AddHeaderWithValue(header string, keyHeader bool, t Type, value string) {
	// TODO: update to check if a key header already exists
	h := &Header{header, keyHeader, t}
	v := &Value{value}
	r.RowMap[h] = v
}

func (r *Row) RemoveHeader(header string) {
	// Ignores if header is Key Header
	newRowMap := map[HeaderI]ValueI{}
	for h, v := range r.RowMap {
		if h.GetName() != header {
			newRowMap[h] = v
		}
	}
}

func (r *Row) UpdateHeaderValue(header string, value string) {
	for h := range r.RowMap {
		if h.GetName() == header {
			r.RowMap[h].SetValue(value)
			break
		}
	}
}
