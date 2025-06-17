package get

import (
	"github.com/Battlekeeper/veyl/internal/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Domain(c *gin.Context) {
	domainId := c.Param("domainid")

	objid, err := primitive.ObjectIDFromHex(domainId)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid domain ID"})
		return
	}

	domain, err := types.GetDomainById(objid)
	if err != nil {
		c.JSON(404, gin.H{"error": "Domain not found"})
		return
	}
	user := c.MustGet("user").(*types.User)
	if domain.Owner != user.Id {
		c.JSON(403, gin.H{"error": "You do not have permission to access this domain"})
		return
	}
	c.JSON(200, domain)
}

func DomainNetworks(c *gin.Context) {
	domainId := c.Param("domainid")

	objid, err := primitive.ObjectIDFromHex(domainId)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid domain ID"})
		return
	}

	domain, err := types.GetDomainById(objid)
	if err != nil {
		c.JSON(404, gin.H{"error": "Domain not found"})
		return
	}
	user := c.MustGet("user").(*types.User)
	if domain.Owner != user.Id {
		c.JSON(403, gin.H{"error": "You do not have permission to access this domain's networks"})
		return
	}

	networks, err := domain.GetNetworks()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve networks"})
		return
	}
	c.JSON(200, networks)
}
