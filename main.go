package main

import (
	"fmt"

	"github.com/brownlow2/pdb/internal/dbmanager"
)

func main() {
	dbm := dbmanager.New()
	headers := []string{
		"Title",
		"Platform",
		"Hours",
		"Plat Name",
		"Number of Trophies",
		"Points Gained",
		"Platinumed",
	}
	dbm.CreateDB("Trophies", headers, "Title")

	row1 := map[string]string{
		"Title":              "Assassin's Creed 2",
		"Platform":           "PS4",
		"Hours":              "35",
		"Plat Name":          "Master Assassin",
		"Number of Trophies": "54",
		"Points Gained":      "1500",
		"Platinumed":         "true",
	}
	err := dbm.DBs["Trophies"].AddRow(row1)
	if err != nil {
		panic(err)
	}

	row2 := map[string]string{
		"Title":              "The Simpsons: Hit and Run",
		"Platform":           "PS5",
		"Hours":              "20",
		"Plat Name":          "D'Oh!",
		"Number of Trophies": "45",
		"Points Gained":      "1500",
		"Platinumed":         "true",
	}
	err = dbm.DBs["Trophies"].AddRow(row2)
	if err != nil {
		panic(err)
	}

	for _, db := range dbm.DBs {
		for _, row := range db.GetRows() {
			for h, v := range row {
				fmt.Printf("%s: %s\n", h, v)
			}
			fmt.Println("---------------------------")
		}
	}
}
