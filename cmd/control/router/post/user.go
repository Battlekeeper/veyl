package post

import (
	"fmt"

	"github.com/Battlekeeper/veyl/internal/types"
	"github.com/gin-gonic/gin"
)

func UserSignup(c *gin.Context) {
	type UserSignupRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	var req UserSignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Check if user already exists
	existingUser, err := types.GetUserByEmail(req.Email)
	if err == nil && existingUser != nil {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}

	user, err := types.CreateUser(req.Email, req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	token, err := types.GenerateJWT(user.Id.Hex())
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func UserLogin(c *gin.Context) {
	type UserLoginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	var req UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := types.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := types.GenerateJWT(user.Id.Hex())
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
