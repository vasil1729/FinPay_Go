// api-gateway/controllers/exampleController.go
package controllers

import "github.com/gin-gonic/gin"

func HelloWorld(c *gin.Context) {
	// print to console
	println("Hello, World!")

	c.JSON(200, gin.H{"message": "Hello, Anha!"})
}
