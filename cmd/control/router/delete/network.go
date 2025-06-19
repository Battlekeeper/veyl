package delete

import (
	"github.com/Battlekeeper/veyl/internal/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Network(c *gin.Context) {
	type NetworkDeleteRequest struct {
		NetworkId string `json:"network_id" binding:"required"`
	}

	var req NetworkDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := c.MustGet("user").(*types.User)

	objid, err := primitive.ObjectIDFromHex(req.NetworkId)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid network ID"})
		return
	}

	network, err := types.GetNetworkById(objid)
	if err != nil {
		c.JSON(404, gin.H{"error": "Network not found"})
		return
	}

	if network.Owner != user.Id {
		c.JSON(403, gin.H{"error": "You do not have permission to delete this network"})
		return
	}

	err = types.DeleteNetwork(objid)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete network"})
		return
	}

	c.JSON(200, gin.H{"message": "Network deleted successfully"})
}
