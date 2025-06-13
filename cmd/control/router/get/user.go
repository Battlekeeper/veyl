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
