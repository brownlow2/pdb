package db

func (h *Header) GetName() string {
	return h.Name
}

func (h *Header) IsString() bool {
	return h.Type == VALUE_STRING
}

func (h *Header) IsNumber() bool {
	return h.Type == VALUE_NUMBER
}

func (h *Header) IsKeyHeader() bool {
	return h.KeyHeader
}
