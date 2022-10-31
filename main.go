package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	// url:http://localhost:3000/employees/sam/role/33
	router.GET("/employees/:username/*rest", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"username": c.Param("username"),
			"rest":     c.Param("rest"),
		})
	})

	log.Fatal(router.Run(":3000"))
}
