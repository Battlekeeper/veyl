package middleware

import (
	"github.com/Battlekeeper/veyl/internal/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		//remove "Bearer " prefix if it exists
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		userid, err := types.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		objid, err := primitive.ObjectIDFromHex(userid)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid user ID format"})
			c.Abort()
			return
		}

		user, err := types.GetUserById(objid)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve user"})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
