package app

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rochimfn/sv-be/schema"
	"gorm.io/gorm"
)

type ArticleHandler interface {
	NewPost(c *gin.Context)
	DetailPost(c *gin.Context)
	ListPost(c *gin.Context)
	UpdatePost(c *gin.Context)
	DeletePost(c *gin.Context)
}

type articleHandler struct {
	DB *gorm.DB
}

func NewArticleHandler(db *gorm.DB) ArticleHandler {
	return &articleHandler{
		DB: db,
	}
}

func (uc *articleHandler) NewPost(c *gin.Context) {
	var postRequest schema.PostRequest
	var newPost schema.Post
	if err := c.ShouldBindJSON(&postRequest); err != nil {
		log.Println("err_bind", err)
		c.JSON(http.StatusBadRequest, schema.GeneralError{
			Error: err.Error(),
		})
		return
	}

	newPost.FromRequest(postRequest)
	r := uc.DB.Create(&newPost)
	if r.Error != nil {
		log.Println("err_bind", r.Error)
		c.JSON(http.StatusExpectationFailed, schema.GeneralError{
			Error: r.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func (uc *articleHandler) DetailPost(c *gin.Context) {
	var post schema.Post
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("err_strconv", err)
		c.JSON(http.StatusBadRequest, schema.GeneralError{
			Error: err.Error(),
		})
		return
	}

	err = uc.DB.Where("Id = ?", id).First(&post).Error
	switch err {
	case nil:
		c.JSON(http.StatusOK, post)
	case gorm.ErrRecordNotFound:
		c.JSON(http.StatusNotFound, schema.GeneralError{
			Error: err.Error(),
		})
	default:
		c.JSON(http.StatusExpectationFailed, schema.GeneralError{
			Error: err.Error(),
		})

	}
}

func (uc *articleHandler) ListPost(c *gin.Context) {
	// rute /article/:limit/:offset tidak bisa digunakan di gin
	// karena akan konflik dengan /article/:id
	// https://github.com/gin-gonic/gin/issues/205
	if c.Param("id") == "" || c.Param("offset") == "" {
		c.JSON(http.StatusBadRequest, schema.GeneralError{
			Error: "limit and offset must be set",
		})
		return
	}

	limit, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("err_strconv", err)
		c.JSON(http.StatusBadRequest, schema.GeneralError{
			Error: err.Error(),
		})
		return
	}

	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		log.Println("err_strconv", err)
		c.JSON(http.StatusBadRequest, schema.GeneralError{
			Error: err.Error(),
		})
		return
	}

	var posts []schema.Post
	r := uc.DB.Limit(limit).Offset(offset).Find(&posts)
	switch r.Error {
	case nil:
		c.JSON(http.StatusOK, posts)
	default:
		c.JSON(http.StatusExpectationFailed, schema.GeneralError{
			Error: err.Error(),
		})
	}
}

func (uc *articleHandler) UpdatePost(c *gin.Context) {
	var postRequest schema.PostRequest
	var updatedPost schema.Post
	if err := c.ShouldBindJSON(&postRequest); err != nil {
		log.Println("err_bind", err)
		c.JSON(http.StatusBadRequest, schema.GeneralError{
			Error: err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("err_strconv", err)
		c.JSON(http.StatusBadRequest, schema.GeneralError{
			Error: err.Error(),
		})
		return
	}

	updatedPost.FromRequest(postRequest)
	err = uc.DB.Where("Id = ?", id).Updates(&updatedPost).Error
	switch err {
	case nil:
		c.JSON(http.StatusAccepted, gin.H{})
	case gorm.ErrRecordNotFound:
		c.JSON(http.StatusNotFound, schema.GeneralError{
			Error: err.Error(),
		})
	default:
		c.JSON(http.StatusExpectationFailed, schema.GeneralError{
			Error: err.Error(),
		})

	}
}
func (uc *articleHandler) DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("err_strconv", err)
		c.JSON(http.StatusBadRequest, schema.GeneralError{
			Error: err.Error(),
		})
		return
	}

	err = uc.DB.Where("Id = ?", id).Delete(&schema.Post{}).Error
	switch err {
	case nil:
		c.JSON(http.StatusNoContent, gin.H{})
	default:
		c.JSON(http.StatusExpectationFailed, schema.GeneralError{
			Error: err.Error(),
		})

	}
}
