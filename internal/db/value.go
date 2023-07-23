package db

func (v *Value) GetValue() string {
	return v.Value
}

func (v *Value) SetValue(value string) {
	v.Value = value
}
