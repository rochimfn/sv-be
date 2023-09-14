package main

import (
	"fmt"
	"log"

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

	handler := app.NewArticleHandler(db)
	r := gin.Default()
	r.POST("/article", handler.NewPost)
	r.GET("/article/:id", handler.DetailPost)
	r.GET("/article/:id/:offset", handler.ListPost)
	// rute /article/:limit/:offset tidak bisa digunakan di gin
	// karena akan konflik dengan /article/:id
	// https://github.com/gin-gonic/gin/issues/205
	r.PUT("/article/:id", handler.UpdatePost)
	r.DELETE("/article/:id", handler.DeletePost)

	log.Println("listen and serve on 0.0.0.0:8080")
	r.Run()
}
