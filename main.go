package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/employee", func(c *gin.Context) {
		c.File("./public/employee.html")
	})

	router.POST("/employee", func(c *gin.Context) {
		c.String(http.StatusOK, "New Request Post!")
	})
	log.Fatal(router.Run(":3000"))
}
