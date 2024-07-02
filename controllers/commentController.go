package controllers

import (
	"BlogManagementSystem/database"
	"BlogManagementSystem/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func GetComments(c *gin.Context) {
	var comments []models.Comment
	if err := database.DB.Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func GetCommentsByUserID(c *gin.Context) {
	UserID := c.Param("user_id")
	var comments []models.Comment
	result := database.DB.Where("user_id = ?", UserID).Find(&comments)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}
func GetCommentsByPostID(c *gin.Context) {
	PostID := c.Param("post_id")
	var comments []models.Comment
	result := database.DB.Where("post_id = ?", PostID).Find(&comments)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
}
func GetCommentByID(c *gin.Context) {
	ID := c.Param("id")
	var comment models.Comment
	result := database.DB.Where("id = ?", ID).First(&comment)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, comment)
}
func CreateComment(c *gin.Context) {
	var newComment models.Comment
	if err := c.BindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	newComment.CreatedAt = time.Now()

	result := database.DB.Create(&newComment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Comment Created", "comment": newComment})
}
func DeleteComment(c *gin.Context) {
	ID := c.Param("id")
	result := database.DB.Delete(&models.Comment{}, ID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment Deleted"})
}
