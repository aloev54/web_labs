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

type CartItem struct {
	ID         int     `json:"id"`
	RecordID   int     `json:"record_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

func AddToCart(db *sql.DB, item CartItem) error {
	query := `INSERT INTO cart_items (record_id, quantity, total_price) VALUES (?,?,?)`
	_, err := db.Exec(query, item.RecordID, item.Quantity, item.TotalPrice)
	return err
}

func GetCartItems(db *sql.DB) ([]CartItem, error) {
	rows, err := db.Query("SELECT id, record_id, quantity, total_price FROM cart_items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cartItems []CartItem
	for rows.Next() {
		var item CartItem
		if err := rows.Scan(&item.ID, &item.RecordID, &item.Quantity, &item.TotalPrice); err != nil {
			return nil, err
		}
		cartItems = append(cartItems, item)
	}
	return cartItems, nil
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
