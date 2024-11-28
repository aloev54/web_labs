package controllers

import (
	"database/sql"
	"lab_4/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

// HomePage обрабатывает запросы к главной странице
func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func MethodOverrideMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {
			if overrideMethod := c.PostForm("_method"); overrideMethod != "" {
				c.Request.Method = overrideMethod
			}
		}
		c.Next()
	}
}

// @Summary Получить все записи
// @Description Возвращает массив элементов
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Record
// @Router /records [get]
func GetRecords(c *gin.Context) {
	records, err := models.GetRecordsDB(db)
	if err != nil {
		// c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}
	// c.IndentedJSON(http.StatusOK, records)
	c.HTML(http.StatusOK, "records.html", records)
}

// @Summary Добавить новую запись
// @Description Добавляет новый элемент
// @Accept  json
// @Produce  json
// @Success 201 {array} models.Record
// @Router /records [post]
func PostRecords(c *gin.Context) {
	log.Println("/records")
	var newRecord models.Record
	if err := c.ShouldBind(&newRecord); err != nil {
		return
	}
	if err := models.AddRecordDB(db, newRecord); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": err.Error()})
		return
	}
	// c.HTML(http.StatusCreated, "success.html", newRecord)
	c.HTML(http.StatusCreated, "success.html", gin.H{
		"message": "Record added successfully!",
		"record":  newRecord,
	})
}

// @Summary Получить запись по id
// @Description Возвращает элемент
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Record
// @Router /records/:id [get]
func GetRecordById(c *gin.Context) {
	log.Println("records/:id")
	id := c.Param("id")
	log.Println(id)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Invalid record ID"})
	}
	query := `SELECT id, title, artist, genre, price FROM records WHERE id = ?`
	row := db.QueryRow(query, idInt)
	var rec models.Record
	err = row.Scan(&rec.ID, &rec.Title, &rec.Artist, &rec.Genre, &rec.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			c.HTML(http.StatusNoContent, "error.html", gin.H{"message": "record not found"})
			// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "record not found"})
		}
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}
	// c.IndentedJSON(http.StatusOK, rec)
	c.HTML(http.StatusOK, "records.html", []models.Record{rec})
}

// @Summary Удалить запись по id
// @Description Удаляет элемент
// @Accept  json
// @Produce  json
// @Success 204 {array} models.Record
// @Router /records/:id [delete]
func DeleteRecordById(c *gin.Context) {
	// log.Println("records/:id")
	// id := c.Param("id")
	id := c.PostForm("id")
	log.Println("Delete record with ID:", id)
	log.Println(id)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		// c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Invalid record ID"})
	}
	query := `DELETE FROM records WHERE id = ?`
	result, err := db.Exec(query, idInt)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"message": "Record not found"})
	}
	c.HTML(http.StatusOK, "success.html", gin.H{
		"message": "Record deleted successfully!",
	})
}

// @Summary Заменить запись по id
// @Description Заменяет элемент
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Record
// @Router /records/:id [put]
func UpdateRecordById(c *gin.Context) {
	id := c.PostForm("id")
	log.Println("Обновление записи с ID:", id)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Invalid record ID"})
		return
	}

	var newRecord models.Record
	if err := c.ShouldBind(&newRecord); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Invalid form data"})
		return
	}

	query := `UPDATE records SET title = ?, artist = ?, genre = ?, price = ? WHERE id = ?`
	result, err := db.Exec(query, newRecord.Title, newRecord.Artist, newRecord.Genre, newRecord.Price, idInt)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	rowAffected, err := result.RowsAffected()
	if err != nil || rowAffected == 0 {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"message": "Record not found"})
		return
	}

	c.HTML(http.StatusOK, "success.html", gin.H{
		"message": "Record updated successfully!",
		"record":  newRecord,
	})
}
