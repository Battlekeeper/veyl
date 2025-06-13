package router

import (
	"fmt"

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

	// POST
	r.POST("/api/user/signup", post.UserSignup)

	err := r.Run(fmt.Sprintf("0.0.0.0:%d", 8080))
	if err != nil {
		panic(fmt.Sprintf("Failed to start router: %v", err))
	}
}
