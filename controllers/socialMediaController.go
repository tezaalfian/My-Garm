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

// CreateSocialMedia creates a new social media
// @Summary creates a new social media
// @Description creates a new social media
// @Tags socialMedia
// @Accept  json
// @Produce  json
// @Param socialMedia body models.SocialMedia true "SocialMedia object"
// @Success 201 {object} models.SocialMedia
// @Failure 400 {object} string
// @Router /social-media [post]
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

// UpdateSocialMedia updates a social media
// @Summary updates a social media
// @Description updates a social media
// @Tags socialMedia
// @Accept  json
// @Produce  json
// @Param socialMediaId path int true "SocialMedia ID"
// @Param socialMedia body models.SocialMedia true "SocialMedia object"
// @Success 200 {object} models.SocialMedia
// @Failure 400 {object} string
// @Router /social-media/{socialMediaId} [put]
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

// DeleteSocialMedia deletes a social media
// @Summary deletes a social media
// @Description deletes a social media
// @Tags socialMedia
// @Accept  json
// @Produce  json
// @Param socialMediaId path int true "SocialMedia ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /social-media/{socialMediaId} [delete]
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

// GetSocialMedia gets a social media
// @Summary gets a social media
// @Description gets a social media
// @Tags socialMedia
// @Accept  json
// @Produce  json
// @Param socialMediaId path int true "SocialMedia ID"
// @Success 200 {object} models.SocialMedia
// @Failure 400 {object} string
// @Router /social-media/{socialMediaId} [get]
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

// GetAllSocialMedias gets all social media
// @Summary gets all social media
// @Description gets all social media
// @Tags socialMedia
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.SocialMedia
// @Failure 400 {object} string
// @Router /social-media [get]
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
