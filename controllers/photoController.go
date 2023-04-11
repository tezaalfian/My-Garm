package controllers

import (
	"MyGarm/database"
	"MyGarm/helpers"
	"MyGarm/models"
	"MyGarm/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
)

// CreatePhoto creates a new photo
// @Summary create a new photo
// @Description create a new photo
// @Tags photos
// @Accept  json
// @Produce  json
// @Param photo body models.Photo true "Photo object"
// @Success 201 {object} models.Photo
// @Failure 400 {object} string
// @Router /photos [post]
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	photoFile, _, err := c.Request.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: photoFile})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.PhotoUrl = uploadUrl
	errCreate := db.Debug().Create(&Photo).Error
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": errCreate.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Photo)
}

// UpdatePhoto updates a photo
// @Summary updates a photo
// @Description updates a photo
// @Tags photos
// @Accept  json
// @Produce  json
// @Param photoId path int true "Photo ID"
// @Param photo body models.Photo true "Photo object"
// @Success 200 {object} models.Photo
// @Failure 400 {object} string
// @Router /photos/{photoId} [put]
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	photoFile, _, err := c.Request.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: photoFile})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	Photo := models.Photo{}

	PhotoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(PhotoId)

	errUpdate := db.Model(&Photo).Where("id = ?", PhotoId).Updates(models.Photo{Title: Photo.Title, PhotoUrl: uploadUrl, Caption: Photo.Caption}).Error

	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": errUpdate.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Photo)
}

// DeletePhoto deletes a photo
// @Summary deletes a photo
// @Description deletes a photo
// @Tags photos
// @Accept  json
// @Produce  json
// @Param photoId path int true "Photo ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /photos/{photoId} [delete]
func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}

	PhotoId, _ := strconv.Atoi(c.Param("photoId"))
	Photo.ID = uint(PhotoId)

	err := db.Debug().Delete(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Photo deleted successfully",
	})
}

// GetPhoto gets a photo
// @Summary gets a photo
// @Description gets a photo
// @Tags photos
// @Accept  json
// @Produce  json
// @Param photoId path int true "Photo ID"
// @Success 200 {object} models.Photo
// @Failure 400 {object} string
// @Router /photos/{photoId} [get]
func GetPhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}

	PhotoId, _ := strconv.Atoi(c.Param("photoId"))
	Photo.ID = uint(PhotoId)

	err := db.Debug().Where("id = ?", PhotoId).Take(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Photo)
}

// GetAllPhotos gets all photos
// @Summary gets all photos
// @Description gets all photos
// @Tags photos
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Photo
// @Failure 400 {object} string
// @Router /photos [get]
func GetAllPhotos(c *gin.Context) {
	db := database.GetDB()
	Photos := []models.Photo{}

	err := db.Debug().Find(&Photos).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Photos)
}
