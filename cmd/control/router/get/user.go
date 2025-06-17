package get

import (
	"github.com/Battlekeeper/veyl/internal/types"
	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	user, exists := c.MustGet("user").(*types.User)
	if !exists {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	user.PasswordHash = ""

	c.JSON(200, user)
}

func UserDomains(c *gin.Context) {
	user, exists := c.MustGet("user").(*types.User)
	if !exists {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	domains, err := types.GetDomainsByUserId(user.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve domains"})
		return
	}

	c.JSON(200, domains)
}
