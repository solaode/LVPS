package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Placeholder route
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "LVPS backend"})
    })

    r.Run(":8080")
}
