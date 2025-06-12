package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()

	err := r.Run(fmt.Sprintf("0.0.0.0:%d", 8080))
	if err != nil {
		panic(fmt.Sprintf("Failed to start router: %v", err))
	}
}
