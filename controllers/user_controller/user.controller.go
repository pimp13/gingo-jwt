package user_controller

import (
	"Gin-GORM-Project/database"
	"Gin-GORM-Project/helpers"
	"Gin-GORM-Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(ctx *gin.Context) {
	var users []models.User

	result := database.DB.Select("id", "username", "email").Find(&users)

	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	hashPassword, err := helpers.MakeHash(user.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = hashPassword

	result := database.DB.Create(&user)
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"data": user})
}
