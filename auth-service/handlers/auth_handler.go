package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"smarket/auth-service/database"
	"smarket/auth-service/models"
	"smarket/auth-service/utils"
)

func Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashed),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists or invalid data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}


func Login(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}

	access, _ := utils.GenerateAccessToken(user.Username)
	refresh, _ := utils.GenerateRefreshToken(user.Username)

	user.RefreshToken = refresh
	database.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func Validate(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	// Remove "Bearer " prefix if present
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	username, err := utils.ValidateToken(tokenString, os.Getenv("ACCESS_SECRET"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "token valid", "user": username})
}


func Refresh(c *gin.Context) {
	var body struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing refresh token"})
		return
	}

	username, err := utils.ValidateToken(body.RefreshToken, os.Getenv("REFRESH_SECRET"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	if user.RefreshToken != body.RefreshToken {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token mismatch"})
		return
	}

	newAccess, _ := utils.GenerateAccessToken(user.Username)
	newRefresh, _ := utils.GenerateRefreshToken(user.Username)
	user.RefreshToken = newRefresh
	database.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  newAccess,
		"refresh_token": newRefresh,
	})
}
