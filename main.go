package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.StaticFile("/staticfile", "./static/index.html")

	router.Run(":8080")

}
