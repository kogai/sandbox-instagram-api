package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setup() error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	err = r.Run(":3000")
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
