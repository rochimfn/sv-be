package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rochimfn/sv-be/app"
	"github.com/rochimfn/sv-be/schema"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	conf := app.InitConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_HOST, conf.DB_PORT, conf.DB_SCHEMA)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("gorm_error", err)
	}
	db.AutoMigrate(&schema.Post{})

	e := gin.Default()
	e = app.InitMiddleware(e)
	handler := app.NewArticleHandler(db)
	router := app.InitArticleRoute(e, handler)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "pong"})
	})
	log.Println("listen and serve on 0.0.0.0:8080")
	router.Run()
}
