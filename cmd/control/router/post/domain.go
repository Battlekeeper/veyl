package post

import (
	"github.com/Battlekeeper/veyl/internal/types"
	"github.com/gin-gonic/gin"
)

func DomainCreate(c *gin.Context) {
	type DomainCreateRequest struct {
		Name string `json:"name" binding:"required"`
	}

	var req DomainCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := c.MustGet("user").(*types.User)
	domain := types.CreateDomain(req.Name, user.Id)

	c.JSON(200, domain)
}
