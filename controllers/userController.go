package controllers

import (
	"MyGarm/database"
	"MyGarm/helpers"
	"MyGarm/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	appJSON = "application/json"
)

// Register
// @Summary register a new user
// @Description register a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User object"
// @Success 201 {object} models.User
// @Failure 400 {object} string
// @Router /users/register [post]
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	err := db.Debug().Create(&User).Error
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": User,
	})
}

// Login
// @Summary login a user
// @Description login a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User object"
// @Success 200 {object} models.User
// @Failure 400 {object} string
// @Router /users/login [post]
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email or password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email or password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
