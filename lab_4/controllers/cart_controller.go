package controllers

import (
	"lab_4/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	recordID := c.PostForm("record_id")
	quantity := c.PostForm("quantity")

	recordIDInt, err := strconv.Atoi(recordID)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Invalid record ID"})
		return
	}

	quantityInt, err := strconv.Atoi(quantity)
	if err != nil || quantityInt <= 0 {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Invalid quantity"})
		return
	}

	var price float64
	query := `SELECT price FROM records WHERE id = ?`
	err = db.QueryRow(query, recordIDInt).Scan(&price)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Record not found"})
		return
	}

	totalPrice := float64(quantityInt) * price

	item := models.CartItem{
		RecordID:   recordIDInt,
		Quantity:   quantityInt,
		TotalPrice: totalPrice,
	}

	if err := models.AddToCart(db, item); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Failed to add to cart"})
		return
	}

	c.HTML(http.StatusOK, "success.html", gin.H{
		"message": "Record added to cart successfully!",
	})
}

func ViewCart(c *gin.Context) {
	cartItems, err := models.GetCartItems(db)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Failed to retrieve cart items"})
		return
	}

	var itemsWithDetails []struct {
		Title      string
		Quantity   int
		Price      float64
		TotalPrice float64
		ID         int
	}

	for _, item := range cartItems {
		var title string
		query := `SELECT title FROM records WHERE id = ?`
		err := db.QueryRow(query, item.RecordID).Scan(&title)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Failed to retrieve record details"})
			return
		}

		itemsWithDetails = append(itemsWithDetails, struct {
			Title      string
			Quantity   int
			Price      float64
			TotalPrice float64
			ID         int
		}{
			Title:      title,
			Quantity:   item.Quantity,
			Price:      item.TotalPrice / float64(item.Quantity),
			TotalPrice: item.TotalPrice,
			ID:         item.RecordID,
		})
	}

	c.HTML(http.StatusOK, "cart.html", itemsWithDetails)
}

func DeleteFromCart(c *gin.Context) {
	recordID := c.PostForm("record_id")
	recordIDInt, err := strconv.Atoi(recordID)
	if err != nil || recordIDInt <= 0 {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Invalid record ID"})
		return
	}

	err = models.DeleteFromCart(db, recordIDInt)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Failed to remove item from cart"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/cart")
}
