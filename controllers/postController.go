package controllers

import (
	"BlogManagementSystem/database"
	"BlogManagementSystem/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post

	if err := database.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func GetPostsByUserID(c *gin.Context) {
	UserID := c.Param("user_id")
	var posts []models.Post
	result := database.DB.Where("user_id = ?", UserID).Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func GetPostByID(c *gin.Context) {
	ID := c.Param("id")
	var post models.Post
	result := database.DB.Where("id = ?", ID).First(&post)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}
func CreatePost(c *gin.Context) {
	var newPost models.Post
	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	newPost.CreatedAt = time.Now()
	newPost.UpdatedAt = time.Now()

	result := database.DB.Create(&newPost)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post Created", "post": newPost})
}

func UpdatePost(c *gin.Context) {
	ID := c.Param("id")
	var post models.Post
	result := database.DB.Where("id = ?", ID).First(&post)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	var updatedPost models.Post
	if err := c.BindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	updatedPost.UpdatedAt = time.Now()
	result = database.DB.Model(&post).Updates(updatedPost)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post updated", "post": updatedPost})
}
func DeletePost(c *gin.Context) {
	ID := c.Param("id")
	result := database.DB.Delete(&models.Post{}, ID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post Deleted "})
}
