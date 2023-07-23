package db

/*
	Test for:
		- The correct value is returned
*/
func (v *Value) GetValue() string {
	return v.Value
}

/*
	Test for:
		- The correct value is set
*/
func (v *Value) SetValue(value string) {
	v.Value = value
}
