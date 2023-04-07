package controllers

import (
	"MyGarm/database"
	"MyGarm/helpers"
	"MyGarm/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
)

// CreateComment creates a new comment
// @Summary creates a new comment
// @Description creates a new comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Param comment body models.Comment true "Comment object"
// @Success 201 {object} models.Comment
// @Failure 400 {object} string
// @Router /comments [post]
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	err := db.Debug().Create(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Comment)
}

// UpdateComment updates a comment
// @Summary updates a comment
// @Description updates a comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Param commentId path int true "Comment ID"
// @Param comment body models.Comment true "Comment object"
// @Success 200 {object} models.Comment
// @Failure 400 {object} string
// @Router /comments/{commentId} [put]
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	CommentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(CommentId)

	err := db.Model(&Comment).Where("id = ?", CommentId).Updates(models.Comment{Message: Comment.Message, PhotoID: Comment.PhotoID}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comment)
}

// DeleteComment deletes a comment
// @Summary deletes a comment
// @Description deletes a comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Param commentId path int true "Comment ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /comments/{commentId} [delete]
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}

	CommentId, _ := strconv.Atoi(c.Param("commentId"))
	Comment.ID = uint(CommentId)

	err := db.Debug().Delete(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Comment deleted successfully",
	})
}

// GetComment gets a comment
// @Summary gets a comment
// @Description gets a comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Param commentId path int true "Comment ID"
// @Success 200 {object} models.Comment
// @Failure 400 {object} string
// @Router /comments/{commentId} [get]
func GetComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}

	CommentId, _ := strconv.Atoi(c.Param("commentId"))
	Comment.ID = uint(CommentId)

	err := db.Debug().Where("id = ?", CommentId).Take(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comment)
}

// GetAllComments gets all comments
// @Summary gets all comments
// @Description gets all comments
// @Tags comments
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Comment
// @Failure 400 {object} string
// @Router /comments [get]
func GetAllComments(c *gin.Context) {
	db := database.GetDB()
	Comments := []models.Comment{}

	err := db.Debug().Find(&Comments).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comments)
}
