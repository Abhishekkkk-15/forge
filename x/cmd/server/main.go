package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "8080"

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"service": "x",
		})
	})

	log.Printf("x running on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
