package some

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// TestGinServer ...
func TestGinServer() {
	fmt.Println("-----TestGinServer-----")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
	fmt.Println("Gin listen and serve on 0.0.0.0:8080")
}
