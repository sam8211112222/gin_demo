package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	router := gin.Default()

	// url: http://localhost:3000/query/?username=sam&year=2022&month=1&month=2&month=3
	router.GET("/query/*rest", func(c *gin.Context) {
		username := c.Query("username")
		year := c.DefaultQuery("year", strconv.Itoa(time.Now().Year()))
		months := c.QueryArray("month")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"year":     year,
			"months":   months,
		})
	})

	router.GET("/employee", func(c *gin.Context) {
		c.File("./public/employee.html")
	})

	// url: http://localhost:3000/employee, try choose a day and amount
	router.POST("/employee", func(c *gin.Context) {
		date := c.PostForm("date")
		amount := c.PostForm("amount")
		username := c.DefaultPostForm("username", "adent")

		c.IndentedJSON(http.StatusOK, gin.H{
			"date":     date,
			"amount":   amount,
			"username": username,
		})

	})

	log.Fatal(router.Run(":3000"))
}
