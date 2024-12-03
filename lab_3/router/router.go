package router

import (
	"lab_3/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Получить все записи
// @Description Возвращает массив элементов
// @Accept  json
// @Produce  json
// @Success 200 {array} data.Record
// @Router /records [get]
func GetRecords(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Records)
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
	for _, record := range data.Records {
		if record.ID == newRecord.ID {
			c.JSON(http.StatusConflict, gin.H{"error": "Record with this ID is already exists"})
			return
		}
	}
	data.Records = append(data.Records, newRecord)
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
	for _, rec := range data.Records {
		if rec.ID == idInt {
			c.IndentedJSON(http.StatusOK, rec)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "record not found"})
}

func DeleteRecordById(c *gin.Context) {
	log.Println("records/:id")
	id := c.Param("id")
	log.Println(id)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	for i, rec := range data.Records {
		if rec.ID == idInt {
			data.Records = append(data.Records[:i], data.Records[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, gin.H{"message": "record deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "record not found"})
}

func UpdateRecordById(c *gin.Context) {
	log.Println("records/:id")
	id := c.Param("id")
	log.Println(id)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	for i, rec := range data.Records {
		if rec.ID == idInt {
			c.BindJSON(&rec)
			data.Records[i] = rec
			c.IndentedJSON(http.StatusOK, rec)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "record not found"})
}
