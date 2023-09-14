package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	public := router.Group("/public")
	public.Use(AuthMiddleware())
	{
		public.StaticFile("/staticfile", "./static/index.html")
	}

	pov := router.Group("/pov")
	{
		pov.GET("/info", func(c *gin.Context) {
			c.String(http.StatusOK, "Public information checklist one.!")
		})
	}
	router.Run(":8080")

}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		duration := time.Since(start)

		connection := c.GetHeader("Connection")
		if connection == "" {
			c.AbortWithStatusJSON(401, gin.H{"Error message": "Unathorized, check connections and permissions"})
		}
		log.Printf("Request - Method: %s | Status: %v | Duration: %d", c.Request.Method, c.Writer.Status(), duration)
		c.Next()
	}

}
