// @title My Project API
// @version 1.0
// @description Это API для работы с записями.
// @host localhost:8080
// @BasePath /

package main

import (
	"lab_4/router"
	"log"

	"database/sql"
	_ "lab_4/docs"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db, err := sql.Open("sqlite3", "./record.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS records (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        artist TEXT,
        genre TEXT,
        price REAL
    );`
	db.Exec(createTableQuery)

	router.SetDB(db)

	r := gin.Default()
	r.GET("/records", router.GetRecords)
	r.POST("/records", router.PostRecords)
	r.GET("/records/:id", router.GetRecordById)
	r.DELETE("/records/:id", router.DeleteRecordById)
	r.PUT("/records/:id", router.UpdateRecordById)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run("localhost:8080")
}
