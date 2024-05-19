package controllers

import (
	"backend-api/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}

	return "Unknown error"
}

func FindPost(c *gin.Context) {
	// get data from database
	var posts []models.Post
	models.DB.Find(&posts)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Lists Data Posts",
		"data":    &posts,
	})
}

func StorePost(c *gin.Context) {
	// validate input
	var input ValidatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
	}
	models.DB.Create(&post)

	c.JSON(201, gin.H{
		"success": true,
		"message": "Post Created Successfully",
		"data":    post,
	})
}

func FindPostById(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Data Post Not Found",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Detail Data Post By ID : " + c.Param("id"),
		"data":    post,
	})
}
