package main

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

func FirstMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Send a response from the middleware
		c.JSON(200, gin.H{
			"message": "First middleware",
		})

		// Read the request body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			println(err)
			println("Error reading request body")
		}
		// Create a map to store the JSON data
		var data map[string]interface{}

		// Unmarshal the JSON data into the map
		err = json.Unmarshal([]byte(string(body)), &data)
		if err != nil {
			println(string(body))
			println(err.Error())
			println("Error parsing JSON data")
			c.AbortWithStatusJSON(400, gin.H{"error": "Error parsing JSON data"})
			return
		}
		println(string(body))
		// You can do further processing with the JSON data here if needed

		// Call the next middleware or handler in the chain
		c.Next()

		// This part of the code will execute after the next handler is called
		c.JSON(200, gin.H{
			"message": "First middleware - after next",
		})
	}
}

	func main() {
		// Create a new Gin router
		router := gin.Default()
		router.Use(FirstMiddleware())
		// Define a route handler
		router.POST("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})
		// Start the server
		router.Run(":8080")
	}
