package ridershipDB

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteRidershipDB struct {
	db *sql.DB
}

func (s *SqliteRidershipDB) Open(filePath string) error {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *SqliteRidershipDB) GetRidership(lineId string) ([]int64, error) {
	query := `
		SELECT SUM(total_ons)
		FROM rail_ridership
		WHERE season = 'Fall 2017'
			AND time_period_id NOT IN ('time_period_10', 'time_period_11')
			AND line_id = ?
		GROUP BY time_period_id
		ORDER BY time_period_id;
	`

	rows, err := s.db.Query(query, lineId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var values []int64
	for rows.Next() {
		var value int64
		err := rows.Scan(&value)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return values, nil
}

func (s *SqliteRidershipDB) Close() error {
	return s.db.Close()
}
