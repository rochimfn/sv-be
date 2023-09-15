package app

import "github.com/gin-gonic/gin"

func InitArticleRoute(r *gin.Engine, handler ArticleHandler) *gin.Engine {
	r.POST("/article", handler.NewPost)
	r.GET("/article/:id", handler.DetailPost)
	r.GET("/article/:id/:offset", handler.ListPost)
	// rute /article/:limit/:offset tidak bisa digunakan di gin
	// karena akan konflik dengan /article/:id
	// https://github.com/gin-gonic/gin/issues/205
	r.PUT("/article/:id", handler.UpdatePost)
	r.DELETE("/article/:id", handler.DeletePost)

	return r
}
