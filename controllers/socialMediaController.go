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

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	err := db.Debug().Create(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, SocialMedia)
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	SocialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(SocialMediaId)

	err := db.Model(&SocialMedia).Where("id = ?", SocialMediaId).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, SocialMedia)
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}

	SocialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	SocialMedia.ID = uint(SocialMediaId)

	err := db.Debug().Delete(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "SocialMedia deleted successfully",
	})
}

func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}

	SocialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	SocialMedia.ID = uint(SocialMediaId)

	err := db.Debug().Where("id = ?", SocialMediaId).Take(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, SocialMedia)
}

func GetAllSocialMedias(c *gin.Context) {
	db := database.GetDB()
	SocialMedias := []models.SocialMedia{}

	err := db.Debug().Find(&SocialMedias).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, SocialMedias)
}
