package router

import (
	"database/sql"
	"lab_4/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

// @Summary Получить все записи
// @Description Возвращает массив элементов
// @Accept  json
// @Produce  json
// @Success 200 {array} data.Record
// @Router /records [get]
func GetRecords(c *gin.Context) {
	records, err := data.GetRecordsDB(db)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, records)
}

// @Summary Добавить новую запись
// @Description Добавляет новый элемент
// @Accept  json
// @Produce  json
// @Success 201 {array} data.Record
// @Router /records [post]
func PostRecords(c *gin.Context) {
	log.Println("/records")
	var newRecord data.Record
	if err := c.BindJSON(&newRecord); err != nil {
		return
	}
	if err := data.AddRecordDB(db, newRecord); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newRecord)
}

func GetRecordById(c *gin.Context) {
	log.Println("records/:id")
	id := c.Param("id")
	log.Println(id)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	query := `SELECT id, title, artist, genre, price FROM records WHERE id = ?`
	row := db.QueryRow(query, idInt)
	var rec data.Record
	err = row.Scan(&rec.ID, &rec.Title, &rec.Artist, &rec.Genre, &rec.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "record not found"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, rec)
}

func DeleteRecordById(c *gin.Context) {
	log.Println("records/:id")
	id := c.Param("id")
	log.Println(id)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	query := `DELETE FROM records WHERE id = ?`
	result, err := db.Exec(query, idInt)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "record not found"})
	}
	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "record deleted"})
}

func UpdateRecordById(c *gin.Context) {
	log.Println("records/:id")
	id := c.Param("id")
	log.Println(id)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	var newRecord data.Record
	if err := c.BindJSON(&newRecord); err != nil {
		return
	}
	query := `UPDATE records SET title = ?, artist = ?, genre = ?, price = ? WHERE id = ?`
	result, err := db.Exec(query, newRecord.Title, newRecord.Artist, newRecord.Genre, newRecord.Price, idInt)
	rowAffected, err := result.RowsAffected()
	if rowAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "record not found"})
	}
	c.IndentedJSON(http.StatusOK, newRecord)
}
