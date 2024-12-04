// @title My Project API
// @version 1.0
// @description Это API для работы с записями.
// @host localhost:8080
// @BasePath /

/* LAB 3 DONE*/

package main

import (
	"lab_3/router"

	_ "lab_3/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	r.GET("/records", router.GetRecords)
	r.POST("/records", router.PostRecords)
	r.GET("/records/:id", router.GetRecordById)
	r.DELETE("/records/:id", router.DeleteRecordById)
	r.PUT("/records/:id", router.UpdateRecordById)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // to run swag -> http://localhost:8080/swagger/index.html
	r.Run("localhost:8080")
}
