package ridershipDB

import (
	// rdb "main/ridership_db"
	"testing"
)

func TestRidershipDBsMatch(t *testing.T) {
	lineIds := []string{"red", "green", "blue", "orange"}

	for _, lineId := range lineIds {
		t.Run(lineId, func(t *testing.T) {
			// Create instances of CsvRidershipDB and SqliteRidershipDB
			var csvDB RidershipDB = &CsvRidershipDB{}       // CSV implementation
			var sqliteDB RidershipDB = &SqliteRidershipDB{} // SQLite implementation

			// Open database connections
			err := csvDB.Open("../mbta.csv")
			if err != nil {
				t.Fatal(err)
			}
			defer csvDB.Close()

			err = sqliteDB.Open("../mbta.sqlite")
			if err != nil {
				t.Fatal(err)
			}
			defer sqliteDB.Close()

			// Retrieve data from both implementations
			csvData, err := csvDB.GetRidership(lineId)
			if err != nil {
				t.Fatal(err)
			}

			sqliteData, err := sqliteDB.GetRidership(lineId)
			if err != nil {
				t.Fatal(err)
			}

			// Compare the data
			if len(csvData) != len(sqliteData) {
				t.Errorf("Lengths of data arrays do not match")
			}

			for i := 0; i < len(csvData); i++ {
				if csvData[i] != sqliteData[i] {
					t.Errorf("Mismatched data at index %d: Csv: %d, SQLite: %d", i, csvData[i], sqliteData[i])
				}
			}
		})
	}
}
