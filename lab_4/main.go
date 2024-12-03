// @title Records API
// @version 1.0
// @description Это API для работы с записями.
// @host localhost:8080
// @BasePath /

package main

import (
	"log"

	"database/sql"
	"lab_4/controllers"
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

	createCartItemsTable := `
	CREATE TABLE IF NOT EXISTS cart_items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		record_id INTEGER NOT NULL,
		quantity INTEGER NOT NULL,
		total_price REAL NOT NULL,
		FOREIGN KEY (record_id) REFERENCES records(id)
	);
	`
	db.Exec(createCartItemsTable)

	controllers.SetDB(db)

	r := gin.Default()

	r.LoadHTMLGlob("views/*")
	r.Static("/static", "./static")

	r.GET("/", controllers.MainPage)
	r.POST("/records", controllers.PostRecords)
	r.GET("/records/:id", controllers.GetRecordById)
	r.GET("/records", controllers.GetRecords)
	r.POST("/records/delete", controllers.DeleteRecordById)
	r.POST("/records/update", controllers.UpdateRecordById)
	r.Use(controllers.MethodOverrideMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //to run swag -> http://localhost:8080/swagger/index.html
	r.POST("/cart/add", controllers.AddToCart)
	r.GET("/cart", controllers.ViewCart)
	r.POST("/remove-from-cart", controllers.DeleteFromCart)
	r.Run("localhost:8080")
}
