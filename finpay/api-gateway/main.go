// api-gateway/main.go
package main

import (
	"finpay/api-gateway/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Add routes and middleware here

	// Define a route
	router.GET("/hello", controllers.HelloWorld)

	router.Run(":8080") // Start the server
}
