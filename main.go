package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	// go to http://localhost:3000/rest and find result on the page
	router.GET("/rest", func(c *gin.Context) {
		url := c.Request.URL.String()
		headers := c.Request.Header
		cookies := c.Request.Cookies()
		c.IndentedJSON(http.StatusOK, gin.H{
			"url":     url,
			"headers": headers,
			"cookies": cookies,
		})
	})
	log.Fatal(router.Run(":3000"))
}
