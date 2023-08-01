package db

func newDBWithValues() (*DBImpl, []RowI) {
	rows := &Rows{
		Items: []RowI{
			&Row{
				RowMap: map[HeaderI]ValueI{
					&Header{"Title", true, VALUE_STRING}:  &Value{"test"},
					&Header{"Value", false, VALUE_STRING}: &Value{"test2"},
				},
			},
		},
	}
	db := &DBImpl{
		Name:      "test",
		KeyHeader: "Title",
		Headers: map[HeaderI]struct{}{
			&Header{"Title", true, VALUE_STRING}:  struct{}{},
			&Header{"Value", false, VALUE_STRING}: struct{}{},
		},
		Rows: rows,
	}
	return db, rows.Items
}

func createRow() (RowI, map[HeaderI]ValueI) {
	rowMap := map[HeaderI]ValueI{
		&Header{"Key", true, VALUE_STRING}:     &Value{"key value"},
		&Header{"NotKey", false, VALUE_NUMBER}: &Value{""},
	}

	return &Row{
		rowMap,
	}, rowMap
}

func createRows() (RowsI, RowI, map[HeaderI]ValueI) {
	row, rowMap := createRow()
	rows := &Rows{
		Items: []RowI{row},
	}

	return rows, row, rowMap
}
