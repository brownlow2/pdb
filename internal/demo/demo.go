package demo

import (
	"fmt"

	"github.com/brownlow2/pdb/internal/db"
	"github.com/brownlow2/pdb/internal/dbmanager"
)

type template struct {
	Name      string
	KeyHeader string
	Headers   []db.HeaderI
}

var platTemplate template = template{
	Name:      "Platinum Tracker",
	KeyHeader: "Title",
	Headers: []db.HeaderI{
		&db.Header{"Title", true, db.VALUE_STRING},
		&db.Header{"Platform", false, db.VALUE_STRING},
		&db.Header{"Hours to Platinum", false, db.VALUE_NUMBER},
		&db.Header{"Platinum Name", false, db.VALUE_STRING},
		&db.Header{"Number of Trophies", false, db.VALUE_NUMBER},
		&db.Header{"Points Gained", false, db.VALUE_NUMBER},
		&db.Header{"Platinumed", false, db.VALUE_STRING},
	},
}

var rows []map[string]string = []map[string]string{
	{
		"Title":              "Jak 2",
		"Platform":           "PS4",
		"Hours to Platinum":  "23",
		"Platinum Name":      "Done Done Done",
		"Number of Trophies": "41",
		"Points Gained":      "1500",
		"Platinumed":         "1",
	},
	{
		"Title":              "Hogwarts Legacy",
		"Platform":           "PS5",
		"Hours to Platinum":  "55",
		"Platinum Name":      "Trophy Triumph",
		"Number of Trophies": "46",
		"Points Gained":      "1500",
		"Platinumed":         "1",
	},
	{
		"Title":              "Destroy All Humans! 2 Reprobed",
		"Platform":           "PS5",
		"Hours to Platinum":  "",
		"Platinum Name":      "Beyond Perfection",
		"Number of Trophies": "45",
		"Points Gained":      "1500",
		"Platinumed":         "0",
	},
}

func Demo() {
	dbm := dbmanager.New()
	dbm.CreateDB(platTemplate.Name, platTemplate.Headers, platTemplate.KeyHeader)
	platDB, _ := dbm.RetrieveDB("Platinum Tracker")

	for _, row := range rows {
		r := createRowFromMap(row, platTemplate)
		if r == nil {
			panic("row returned empty")
		}
		platDB.AddRow(r)
	}

	fmt.Println("DB Platinum Tracker successfully created")
	fmt.Println("Headers:")
	for _, h := range platDB.GetHeadersString() {
		fmt.Println(h)
	}
	fmt.Println()

	printRows(platDB)
	platDB.AddValueToHeader("20", "Hours to Platinum", "Destroy All Humans! 2 Reprobed")
	platDB.AddValueToHeader("1", "Platinumed", "Destroy All Humans! 2 Reprobed")
	printRows(platDB)
}

func printRows(d db.DB) {
	fmt.Println("Rows:")
	for _, r := range d.GetRows() {
		rowMap := r.GetRowMap()
		for h, v := range rowMap {
			fmt.Printf("%s: %s\n", h.GetName(), v.GetValue())
		}
		fmt.Println()
	}
}

func createRowFromMap(m map[string]string, t template) db.RowI {
	rowMap := make(map[db.HeaderI]db.ValueI, 0)
	row := &db.Row{rowMap}
	for h, v := range m {
		header := getHeader(h, t.Headers)
		if header == nil {
			return &db.Row{}
		}
		row.AddHeaderWithValue(h, header.IsKeyHeader(), header.GetType(), v)
	}
	return row
}

func getHeader(h string, headers []db.HeaderI) db.HeaderI {
	for _, header := range headers {
		if header.GetName() == h {
			return header
		}
	}
	return nil
}
