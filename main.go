package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func setup() error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/on_done", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "done"})
	})

	r.GET("/on_deauth", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "deauth"})
	})

	r.GET("/on_delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "delete"})
	})

	err := r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := setup()
	if err != nil {
		panic(err)
	}
}
