package user_controller

import (
	"Gin-GORM-Project/database"
	"Gin-GORM-Project/helpers"
	"Gin-GORM-Project/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate *validator.Validate

func uniqueEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	var user models.User
	result := database.DB.Select("email").Where("email = ?", email).First(&user)

	return result.Error != nil
}

func emailExists(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	var user models.User
	result := database.DB.Select("email").Where("email = ?", email).First(&user)

	return result.RowsAffected > 0
}

func Register(ctx *gin.Context) {
	// Get email and password and username from request body
	var body struct {
		Email    string `json:"email" validate:"required,email,uniqueEmail"`
		Password string `json:"password" validate:"required,min=8,max=190"`
		Username string `json:"username" validate:"required,min=3,max=125"`
	}
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Failed to bind body: %s", err.Error()),
		})
		return
	}

	// Validate the input
	validate = validator.New()
	validate.RegisterValidation("uniqueEmail", uniqueEmail)

	if err := validate.Struct(body); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Validation error: %s - %s", e.Field(), e.Error()),
			})
			return
		}
	}

	// Hash the password
	hashedPassword, err := helpers.MakeHash(body.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to hash password: %s", err.Error()),
		})
		return
	}

	// Create new user object with hashed password
	user := &models.User{
		Email:    body.Email,
		Password: hashedPassword,
		Username: body.Username,
	}

	// Save the user to the database
	if err := database.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to save user to database: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data":    user,
	})
}

func Login(ctx *gin.Context) {
	// Get email and password from request body
	var body struct {
		Email    string
		Password string
	}
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Failed to bind body: %s", err.Error()),
		})
		return
	}

	// Look up requested user
	var user models.User
	database.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Check if password matches hashed password
	if err := helpers.CheckHash(user.Password, body.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Generate JWT token
	tokenString, err := helpers.GenerateJWT(user.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// Set up the cookie for token
	ctx.SetSameSite(http.SameSiteStrictMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24, "", "localhost", false, true)

	// Send it back
	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Validate(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if exists {
		ctx.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
	}
}

func Logout(c *gin.Context) {
	// Find cookie and delete cookie
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("Authorization", "", -1, "", "localhost", false, true)

	// Send response
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}
