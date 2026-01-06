package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "{{.Port}}"

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"service": "{{.ProjectName}}",
		})
	})

	log.Printf("{{.ProjectName}} running on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
