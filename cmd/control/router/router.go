package router

import (
	"fmt"

	"github.com/Battlekeeper/veyl/cmd/control/router/delete"
	"github.com/Battlekeeper/veyl/cmd/control/router/get"
	"github.com/Battlekeeper/veyl/cmd/control/router/middleware"
	"github.com/Battlekeeper/veyl/cmd/control/router/post"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()

	// Set up middleware
	r.Use(gin.Recovery())

	// GET
	r.GET("/api/user", middleware.UserAuthentication(), get.User)
	r.GET("/api/user/domains", middleware.UserAuthentication(), get.UserDomains)

	r.GET("/api/domain/:domainid", middleware.UserAuthentication(), get.Domain)
	r.GET("/api/domain/:domainid/networks", middleware.UserAuthentication(), get.DomainNetworks)

	r.GET("/api/network/:networkid", middleware.UserAuthentication(), get.Network)

	// POST
	r.POST("/api/user/signup", post.UserSignup)
	r.POST("/api/user/login", post.UserLogin)
	r.POST("/api/domain/create", middleware.UserAuthentication(), post.DomainCreate)
	r.POST("/api/network/create", middleware.UserAuthentication(), post.NetworkCreate)

	// DELETE
	r.DELETE("/api/network/:networkid", middleware.UserAuthentication(), delete.Network)
	r.DELETE("/api/domain/:domainid", middleware.UserAuthentication(), delete.Domain)

	err := r.Run(fmt.Sprintf("0.0.0.0:%d", 8080))
	if err != nil {
		panic(fmt.Sprintf("Failed to start router: %v", err))
	}
}
