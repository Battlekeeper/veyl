package post

import (
	"github.com/Battlekeeper/veyl/internal/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NetworkCreate(c *gin.Context) {
	type NetworkCreateRequest struct {
		Name     string `json:"name" binding:"required"`
		DomainId string `json:"domain_id" binding:"required"`
	}

	var req NetworkCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := c.MustGet("user").(*types.User)

	objid, err := primitive.ObjectIDFromHex(req.DomainId)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid domain ID"})
		return
	}
	domain, err := types.GetDomainById(objid)
	if err != nil {
		c.JSON(404, gin.H{"error": "Domain not found"})
		return
	}
	if domain.Owner != user.Id {
		c.JSON(403, gin.H{"error": "You do not have permission to create a network in this domain"})
		return
	}

	network := types.CreateNetwork(req.Name)
	err = domain.AddNetwork(network.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to associate network with domain"})
		return
	}

	c.JSON(200, network)
}
