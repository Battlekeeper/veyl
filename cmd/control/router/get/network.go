package get

import (
	"github.com/Battlekeeper/veyl/internal/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Network(c *gin.Context) {

	networkId := c.Param("networkid")

	objid, err := primitive.ObjectIDFromHex(networkId)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid network ID"})
		return
	}

	network, err := types.GetNetworkById(objid)
	if err != nil {
		c.JSON(404, gin.H{"error": "Network not found"})
		return
	}

	user := c.MustGet("user").(*types.User)
	if network.Owner != user.Id {
		c.JSON(403, gin.H{"error": "You do not have permission to access this network"})
		return
	}

	c.JSON(200, network)
}
