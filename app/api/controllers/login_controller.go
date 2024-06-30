package controllers

import (
	"app/api/models"
	"app/api/services"
	"app/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req loginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	// メールアドレスを検証
	if err := services.GetUserByEmail(req.Email, &user); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// パスワードを検証
	if !utils.CheckPasswordHash(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}