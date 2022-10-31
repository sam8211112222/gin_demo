package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	// http://localhost:3000/admin --> 404 page not found
	adminGroup := router.Group("/admin")
	// http://localhost:3000/admin/users
	adminGroup.GET("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "page to admin users")
	})
	// http://localhost:3000/admin/roles
	adminGroup.GET("/roles", func(c *gin.Context) {
		c.String(http.StatusOK, "page to admin roles")
	})

	// http://localhost:3000/admin/polices
	adminGroup.GET("/polices", func(c *gin.Context) {
		c.String(http.StatusOK, "page to admin polices")
	})
	log.Fatal(router.Run(":3000"))
}
