package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	router.StaticFile("/", "./public/index.html") // http://localhost:3000/

	router.Static("/public", "./public") // http://localhost:3000/public/employee.html

	router.StaticFS("/fs", http.Dir("./public")) // http://localhost:3000/fs/employee.html

	log.Fatal(router.Run(":3000"))
}
