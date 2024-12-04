package models

import "database/sql"

// Record представляет модель данных для пластинки
type Record struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Genre  string  `json:"genre"`
	Price  float64 `json:"price"`
}

func AddRecordDB(db *sql.DB, record Record) error {
	query := `INSERT INTO records (title, artist, genre, price) VALUES (?,?,?,?)`
	_, err := db.Exec(query, record.Title, record.Artist, record.Genre, record.Price)
	return err
}

func GetRecordsDB(db *sql.DB) ([]Record, error) {
	rows, err := db.Query("SELECT id, title, artist, genre, price FROM records")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var records []Record
	for rows.Next() {
		var record Record
		if err := rows.Scan(&record.ID, &record.Title, &record.Artist, &record.Genre, &record.Price); err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}
