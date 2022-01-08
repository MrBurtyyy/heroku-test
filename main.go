package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func handleIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

const PORT = "PORT"

func main() {
	r := gin.Default()
	r.GET("/", handleIndex)
	r.Run("localhost:" + os.Getenv(PORT))
}
