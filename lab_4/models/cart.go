package models

import (
	"database/sql"
)

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

func DeleteFromCart(db *sql.DB, recordID int) error {
	_, err := db.Exec("DELETE FROM cart_items WHERE record_id = ?", recordID)
	return err
}

// DeleteFromCart удаляет товар из корзины по его ID
// func DeleteFromCart(db *sql.DB, recordID int) error {
// 	// Проверяем, существует ли запись в корзине с таким ID
// 	var exists bool
// 	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM cart_items WHERE record_id = ?)", recordID).Scan(&exists)
// 	if err != nil {
// 		return fmt.Errorf("failed to check if item exists: %w", err)
// 	}
// 	if !exists {
// 		return fmt.Errorf("item not found in cart")
// 	}

// 	// Выполняем удаление записи
// 	_, err = db.Exec("DELETE FROM cart_items WHERE record_id = ?", recordID)
// 	if err != nil {
// 		return fmt.Errorf("failed to delete item from cart: %w", err)
// 	}
// 	return nil
// }
