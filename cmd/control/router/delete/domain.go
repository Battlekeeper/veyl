package delete

import (
	"github.com/Battlekeeper/veyl/internal/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Domain(c *gin.Context) {
	type DomainDeleteRequest struct {
		DomainId string `json:"domain_id" binding:"required"`
	}

	var req DomainDeleteRequest
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
		c.JSON(403, gin.H{"error": "You do not have permission to delete this domain"})
		return
	}

	err = types.DeleteDomain(objid)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete domain"})
		return
	}

	c.JSON(200, gin.H{"message": "Domain deleted successfully"})
}
